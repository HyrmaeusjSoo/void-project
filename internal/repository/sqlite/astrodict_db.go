package sqlite

import (
	"errors"
	"void-project/internal/model"
	"void-project/internal/repository/driver"

	"gorm.io/gorm"
)

type AstrodictRepository struct {
	db *gorm.DB
}

func NewAstrodictRepository() *AstrodictRepository {
	return &AstrodictRepository{db: driver.SQLite}
}

// 查询
func (a *AstrodictRepository) GetList(name string) ([]model.Astrodict, error) {
	ad := []model.Astrodict{}
	err := a.db.Where("C LIKE ?", "%"+name+"%").
		Or("E LIKE ?", "%"+name+"%").
		Find(&ad).Error
	return ad, err
}

// 批量新增
//
//	一次性的同步接口用，而且只是词典，所以清空旧的插入新的
//	lang => 指定语言为ce还是ec，对应到中-英 英-中
func (a *AstrodictRepository) Create(lang string, ad []*model.Astrodict) error {
	var m any
	if lang == "ce" {
		m = &model.Astrodict{}
	} else if lang == "ec" {
		m = &model.AstrodictEC{}
	}
	a.db.Where("TRUE").Delete(m)
	tx := a.db.Model(m)

	// 分批存
	err := tx.Transaction(func(tx *gorm.DB) error {
		cursor := len(ad) / 100
		if rema := len(ad) % 100; rema > 0 {
			cursor++
		}
		for i := 0; i < cursor; i++ {
			x := i * 100
			y := (i + 1) * 100
			if i == cursor-1 {
				y = len(ad)
			}
			if rtx := tx.Create(ad[x:y]); rtx.RowsAffected == 0 || rtx.Error != nil {
				return errors.New("保存失败")
			}
		}
		return nil
	})
	return err
}
