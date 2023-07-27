package repository

import (
	"strings"
	"void-project/global"
	"void-project/internal/model/base"
	"void-project/pkg"
	"void-project/pkg/convert"
	"void-project/pkg/necromancy"

	"gorm.io/gorm"
)

// 分页查询
func Paginate(db *gorm.DB, list any, pager base.Pager) (total int64, err error) {
	err = db.Model(list).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Scopes(PageScope(pager)).Find(list).Error
	if err != nil {
		return
	}

	return
}

// 分页查询Scope
func PageScope(pager base.Pager) func(db *gorm.DB) *gorm.DB {
	pager.Page = pkg.IfElse(pager.Page < 1, 1, pager.Page)
	pager.Size = pkg.IfElse(pager.Size < 1, global.Config.System.PageSize, pager.Size)

	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(pager.Size).Offset((pager.Page - 1) * pager.Size)
	}
}

// Cursor方式分页查询Scope
//
//	使用Scope方法要在查询后手动base.Next = Cursor.Encode()
func CursorScope(cursor base.Cursor) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if strings.ToLower(cursor.SortType) == "desc" {
			db = db.Where(cursor.Field+" < ?", cursor.CursorID).Order(cursor.Field + " DESC")
		} else {
			db = db.Where(cursor.Field+" > ?", cursor.CursorID).Order(cursor.Field + " ASC")
		}
		return db.Limit(cursor.Size)
	}
}

// Cursor方式分页查询
//
//	在Cursor内指定字段和排序方式：
//	首次查询指定（使用Cursor.IsEmpty()判断是否为首次查询），之后分页查询自动解码出信息不需要再指定
//
//	Cursor.Field => 指定条件字段，其字段类型必须为可排序的。例："id"， "user.id"，"user.login_time"
//	Cursor.CursorID => 该值为排序字段的具体值。例："123"，"2023-07-27 15:41:46"，"abc"
//	Cursor.SortType => 指定排序方式。例： "DESC"，"ASC"
func CursorPaginate[T any](db *gorm.DB, list *[]T, cursor base.Cursor) (next base.Next, err error) {
	next = base.Next("")
	// 默认分页条数
	if cursor.Size == 0 {
		cursor.Size = global.Config.System.PageSize
	}

	// 查询
	err = db.Model(list).Scopes(CursorScope(cursor)).Find(list).Error
	if err != nil {
		return
	}
	if len(*list) == 0 || len(*list) < cursor.Size {
		return
	}

	if cursor.Field == "" {
		cursor.Field = "id"
	}
	// 以下是反射字段值
	nameSlc := strings.Split(cursor.Field, ".")
	val, err := necromancy.Extraction((*list)[len(*list)-1], convert.SnakeToPascal(nameSlc[len(nameSlc)-1]))
	if err != nil {
		return
	}
	cursor.CursorID = val

	next = cursor.Encode()
	return
}
