package global

import (
	"encoding/json"
	"os"
	"time"
	"void-project/pkg"

	"github.com/redis/go-redis/v9"
)

type databaseConfig struct {
	MySQL struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     uint16 `json:"port"`
		DBName   string `json:"dbname"`
	} `json:"mysql"`

	SQLite struct {
		Path     string `json:"path"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"sqlite"`

	SQLServer struct{} `json:"sqlserver"`
}

type cacheConfig struct {
	Redis redis.Options `json:"redis"`
}

type systemConfig struct {
	AuthTokenExpire      time.Duration `json:"auth_token_expire"`
	AuthJwtSecret        string        `json:"auth_jwt_secret"`
	AstroDictCacheExpire time.Duration `json:"astro_dict_cache_expire"`
	PageSize             int           `json:"pagesize"`
}

type globalConfig struct {
	DB     databaseConfig
	Cache  cacheConfig
	System systemConfig
}

var Config = &globalConfig{}

func InitConfig() {
	root := pkg.GetRootPath() + "/config/"

	dbConfigFile, _ := os.Open(root + "database.json")
	defer dbConfigFile.Close()
	if err := json.NewDecoder(dbConfigFile).Decode(&Config.DB); err != nil {
		panic(err)
	}

	cacheConfigFile, _ := os.Open(root + "cache.json")
	defer cacheConfigFile.Close()
	if err := json.NewDecoder(cacheConfigFile).Decode(&Config.Cache); err != nil {
		panic(err)
	}

	systemConfigFile, _ := os.Open(root + "system.json")
	defer systemConfigFile.Close()
	if err := json.NewDecoder(systemConfigFile).Decode(&Config.System); err != nil {
		panic(err)
	}
}
