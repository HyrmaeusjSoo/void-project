package scope

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 新增
func Create(db *gorm.DB, model any) error {
	if tx := db.Omit(clause.Associations).Create(model); tx.RowsAffected == 0 {
		if tx.Error != nil {
			return tx.Error
		}
		return errors.New("新增失败")
	}
	return nil
}

// 更新
func Update(db *gorm.DB, model any, where ...any) error {
	if tx := db.Updates(model); tx.RowsAffected == 0 {
		if tx.Error != nil {
			return tx.Error
		}
		return errors.New("更新0条记录")
	}
	return nil
}

// 删除
func Delete(db *gorm.DB, model any, conds ...any) error {
	if tx := db.Delete(model, conds); tx.RowsAffected == 0 {
		if tx.Error != nil {
			return tx.Error
		}
		return errors.New("删除0条记录")
	}
	return nil
}
