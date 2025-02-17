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

// 配置文件
type Config struct {
	// 数据库配置
	DB DB

	// 缓存库配置
	Cache Cache

	// 系统配置
	System System
}

// 系统配置
type System struct {
	// 模式：release=发布模式，dev=开发模式。
	// 主要控制日志只写入文件，禁用日志颜色等; 开发模式更方便查看日志输出，发布模式写入到文件更稳定。
	Mode                 string `json:"mode" yaml:"mode"`
	ListenAddr           string `json:"listen_addr" yaml:"listen_addr"`                         // 服务监听的连接地址和端口号，如：127.0.0.1:5555
	AuthJwtSecret        string `json:"auth_jwt_secret" yaml:"auth_jwt_secret"`                 // 鉴权密钥，推荐使用生成的Hash串
	AuthTokenExpire      int    `json:"auth_token_expire" yaml:"auth_token_expire"`             // 鉴权过期时间（小时）
	PageSize             int    `json:"pagesize" yaml:"pagesize"`                               // 分页大小
	AstroDictCacheExpire int    `json:"astro_dict_cache_expire" yaml:"astro_dict_cache_expire"` // 忽略
	PProf                bool   `json:"pprof" yaml:"pprof"`                                     // 是否开启性能分析
	PProfAddr            string `json:"pprof_addr" yaml:"pprof_addr"`                           // 性能分析链接地址
	VolcAccessKey        string `json:"volc_access_key" yaml:"volc_access_key"`                 // 火山翻译AccessKey
	VolcSecretKey        string `json:"volc_secret_key" yaml:"volc_secret_key"`                 // 火山翻译SecretKey
	StorageLocation      string `json:"storage_location" yaml:"storage_location"`               // 文件存储的本地路径
}

// 缓存库配置
type Cache struct {
	// Redis配置 直接使用go-redis的配置字段更方便
	Redis redis.Options `json:"redis" yaml:"redis"`
}

// 数据库配置
type DB struct {
	// MySQL配置
	MySQL MySQL `json:"mysql" yaml:"mysql"`
	// SQLite配置
	SQLite SQLite `json:"sqlite" yaml:"sqlite"`
	// SQLServer配置
	//SQLServer struct{} `json:"sqlserver" yaml:"sqlserver"`
}

// 数据库
type Databaser interface {
	// 获取配置文件数据库连接dsn字符串
	Dsn() string
	// 获取配置文件SQL日志级别
	SQLLogLevel() string
	// 获取配置文件 最大空闲连接数，最大打开连接数，连接保持时间
	ConnsNumber() (maxIdleConns, maxOpenConns, maxLifetime int)
}

// MySQL配置文件参数
type MySQL struct {
	User         string `json:"user" yaml:"user"`                     // 用户
	Password     string `json:"password" yaml:"password"`             // 密码
	Host         string `json:"host" yaml:"host"`                     // 主机地址
	Port         uint16 `json:"port" yaml:"port"`                     // 端口号
	DBName       string `json:"dbname" yaml:"dbname"`                 // 数据库名
	LogLevel     string `json:"log_level" yaml:"log_level"`           // silent,error,warn,info
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` // 空闲连接池中的最大数
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` // 与数据库打开的最大连接数
	MaxLifetime  int    `json:"max_lifetime" yaml:"max_lifetime"`     // 连接保持时间
}

func (m *MySQL) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.DBName)
}

func (m *MySQL) SQLLogLevel() string {
	return m.LogLevel
}

func (m *MySQL) ConnsNumber() (int, int, int) {
	return m.MaxIdleConns, m.MaxOpenConns, m.MaxLifetime
}

// SQLite配置文件参数
type SQLite struct {
	Path         string `json:"path" yaml:"path"`                     // 路径
	User         string `json:"user" yaml:"user"`                     // 用户
	Password     string `json:"password" yaml:"password"`             // 密码
	LogLevel     string `json:"log_level" yaml:"log_level"`           // silent,error,warn,info
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` // 空闲连接池中的最大数
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` // 与数据库打开的最大连接数
	MaxLifetime  int    `json:"max_lifetime" yaml:"max_lifetime"`     // 连接保持时间
}

func (s *SQLite) Dsn() string {
	return pkg.GetRootPath() + s.Path
}

func (s *SQLite) SQLLogLevel() string {
	return s.LogLevel
}

func (m *SQLite) ConnsNumber() (int, int, int) {
	return m.MaxIdleConns, m.MaxOpenConns, m.MaxLifetime
}

// 配置文件
var Configs Config

// 初始化配置
func InitConfig() {
	config := map[string]any{
		"database": &Configs.DB,
		"cache":    &Configs.Cache,
		"system":   &Configs.System,
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
