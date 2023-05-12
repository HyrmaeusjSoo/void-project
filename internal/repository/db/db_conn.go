package db

import "gorm.io/gorm"

var (
	MySQL     *gorm.DB
	SQLite    *gorm.DB
	SQLServer *gorm.DB
)
