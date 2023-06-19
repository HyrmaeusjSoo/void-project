package repository

import (
	"void-project/pkg"

	"gorm.io/gorm"
)

// 分页查询
func Paginate(db *gorm.DB, list any, page, size int) (total int64, err error) {
	page = pkg.IfElse(page < 1, 1, page)
	size = pkg.IfElse(size < 1, 15, size)

	err = nil
	total = 0

	err = db.Model(list).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(size).Offset((page - 1) * size).Find(list).Error
	if err != nil {
		return
	}

	return
}
