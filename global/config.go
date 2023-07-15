package global

import (
	"encoding/json"
	"fmt"
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
	Mode                 string        `json:"mode"`
	ListenAddr           string        `json:"listen_addr"`
	AuthJwtSecret        string        `json:"auth_jwt_secret"`
	AuthTokenExpire      time.Duration `json:"auth_token_expire"`
	AstroDictCacheExpire time.Duration `json:"astro_dict_cache_expire"`
	PageSize             int           `json:"pagesize"`
}

type config struct {
	DB     databaseConfig
	Cache  cacheConfig
	System systemConfig
}

var Config = &config{}

// 读取配置文件
func InitConfig() {
	sepr := string(os.PathSeparator)
	root := fmt.Sprintf("%v%vconfig%v", pkg.GetRootPath(), sepr, sepr)

	dbConfigFile, err := os.Open(root + "database.json")
	if err != nil {
		panic(err)
	}
	defer dbConfigFile.Close()
	if err := json.NewDecoder(dbConfigFile).Decode(&Config.DB); err != nil {
		panic(err)
	}

	cacheConfigFile, err := os.Open(root + "cache.json")
	if err != nil {
		panic(err)
	}
	defer cacheConfigFile.Close()
	if err := json.NewDecoder(cacheConfigFile).Decode(&Config.Cache); err != nil {
		panic(err)
	}

	systemConfigFile, err := os.Open(root + "system.json")
	if err != nil {
		panic(err)
	}
	defer systemConfigFile.Close()
	if err := json.NewDecoder(systemConfigFile).Decode(&Config.System); err != nil {
		panic(err)
	}
}
