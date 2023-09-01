package global

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"void-project/pkg"

	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
)

// 配置
var Config = struct {
	// 数据库配置
	DB struct {
		// MySQL配置
		MySQL struct {
			User     string `json:"user" yaml:"user"`         // 用户
			Password string `json:"password" yaml:"password"` // 密码
			Host     string `json:"host" yaml:"host"`         // 主机地址
			Port     uint16 `json:"port" yaml:"port"`         // 端口号
			DBName   string `json:"dbname" yaml:"dbname"`     // 数据库名
		} `json:"mysql" yaml:"mysql"`
		// SQLite配置
		SQLite struct {
			Path     string `json:"path" yaml:"path"` // 路径
			User     string `json:"user" yaml:"user"`
			Password string `json:"password" yaml:"password"`
		} `json:"sqlite"`
		// SQLServer配置
		SQLServer struct{} `json:"sqlserver" yaml:"sqlserver"`
	}

	// 缓存库配置
	Cache struct {
		// Redis配置 直接使用go-redis的配置字段更方便
		Redis redis.Options `json:"redis" yaml:"redis"`
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
		Mode                 string `json:"mode" yaml:"mode"`
		ListenAddr           string `json:"listen_addr" yaml:"listen_addr"`                         // 服务监听的连接地址和端口号，如：127.0.0.1:5555
		AuthJwtSecret        string `json:"auth_jwt_secret" yaml:"auth_jwt_secret"`                 // 鉴权密钥，推荐使用生成的Hash串
		AuthTokenExpire      int    `json:"auth_token_expire" yaml:"auth_token_expire"`             // 鉴权过期时间
		PageSize             int    `json:"pagesize" yaml:"pagesize"`                               // 分页大小
		AstroDictCacheExpire int    `json:"astro_dict_cache_expire" yaml:"astro_dict_cache_expire"` // 忽略
		PProf                bool   `json:"pprof" yaml:"pprof"`                                     // 是否开启性能分析
		PProfAddr            string `json:"pprof_addr" yaml:"pprof_addr"`                           // 性能分析链接地址
	}
}{}

// 初始化配置
func InitConfig() {
	config := map[string]any{
		"database": &Config.DB,
		"cache":    &Config.Cache,
		"system":   &Config.System,
	}

	root := fmt.Sprintf("%v%vconfig", pkg.GetRootPath(), string(os.PathSeparator))
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(info.Name())
		if ext != ".json" && ext != ".yaml" {
			return nil
		}

		cfgFile, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer cfgFile.Close()

		var decoder interface {
			Decode(v any) error
		}
		if ext == ".json" {
			decoder = json.NewDecoder(cfgFile)
		} else if ext == ".yaml" {
			decoder = yaml.NewDecoder(cfgFile)
		}

		name := strings.Split(info.Name(), ".")[0]
		if opt, ok := config[name]; ok {
			if err := decoder.Decode(opt); err != nil {
				panic(err)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}
