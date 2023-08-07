package global

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"void-project/pkg"

	"github.com/redis/go-redis/v9"
)

// 配置
var Config = struct {
	// 数据库配置
	DB struct {
		// MySQL配置
		MySQL struct {
			User     string `json:"user"`     // 用户
			Password string `json:"password"` // 密码
			Host     string `json:"host"`     // 主机地址
			Port     uint16 `json:"port"`     // 端口号
			DBName   string `json:"dbname"`   // 数据库名
		} `json:"mysql"`
		// SQLite配置
		SQLite struct {
			Path     string `json:"path"` // 路径
			User     string `json:"user"`
			Password string `json:"password"`
		} `json:"sqlite"`
		// SQLServer配置
		SQLServer struct{} `json:"sqlserver"`
	}

	// 缓存库配置
	Cache struct {
		// Redis配置 直接使用go-redis的配置字段更方便
		Redis redis.Options `json:"redis"`
	}

	// 系统配置
	//	Mode                 string        // 模式：release=发布模式，dev=开发模式。
	//	ListenAddr           string        // 服务监听的连接地址和端口号，如：127.0.0.1:80
	//	AuthJwtSecret        string        // 鉴权密钥，推荐使用生成的Hash串
	//	AuthTokenExpire      time.Duration // 鉴权过期时间
	//	AstroDictCacheExpire time.Duration // 忽略
	//	PageSize             int           // 分页大小
	System struct {
		// 模式：release=发布模式，dev=开发模式。
		// 主要控制日志只写入文件，禁用日志颜色等; 开发模式更方便查看日志输出，发布模式写入到文件更稳定。
		Mode                 string        `json:"mode"`
		ListenAddr           string        `json:"listen_addr"`             // 服务监听的连接地址和端口号，如：127.0.0.1:80
		AuthJwtSecret        string        `json:"auth_jwt_secret"`         // 鉴权密钥，推荐使用生成的Hash串
		AuthTokenExpire      time.Duration `json:"auth_token_expire"`       // 鉴权过期时间
		PageSize             int           `json:"pagesize"`                // 分页大小
		AstroDictCacheExpire time.Duration `json:"astro_dict_cache_expire"` // 忽略
	}
}{}

// 初始化配置
func InitConfig() {
	sep := string(os.PathSeparator)
	root := fmt.Sprintf("%v%vconfig%v", pkg.GetRootPath(), sep, sep)

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
