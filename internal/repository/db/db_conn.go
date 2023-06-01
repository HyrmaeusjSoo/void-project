package db

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	MySQL     *gorm.DB
	SQLite    *gorm.DB
	SQLServer *gorm.DB
	Redis     *redis.Client
)
