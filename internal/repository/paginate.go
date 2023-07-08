package repository

import (
	"void-project/global"
	"void-project/internal/model/base"
	"void-project/pkg"

	"gorm.io/gorm"
)

// 分页查询
func Paginate(db *gorm.DB, list any, pager base.Pager) (total int64, err error) {
	pager.Page = pkg.IfElse(pager.Page < 1, 1, pager.Page)
	pager.Size = pkg.IfElse(pager.Size < 1, global.Config.System.PageSize, pager.Size)

	err = db.Model(list).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(pager.Size).Offset((pager.Page - 1) * pager.Size).Find(list).Error
	if err != nil {
		return
	}

	return
}

// 分页Scope
func PageScope(pager base.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(pager.Size).Offset((pager.Page - 1) * pager.Size)
	}
}
