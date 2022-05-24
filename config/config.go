package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type (
	// Config 配置文件结构体
	Config struct {
		vp *viper.Viper
		*ServerConfig
		*LogConfig
		*MySQLConfig
		*BaseConfig
		*EtcdConfig
		*DurationConfig
	}

	// BaseConfig 基础配置
	BaseConfig struct {
		Host      string `mapstructure:"host"`
		Name      string `mapstructure:"name"`
		Mode      string `mapstructure:"mode"`
		Version   string `mapstructure:"version"`
		StartTime string `mapstructure:"start_time"`
		Port      string `mapstructure:"port"`
		MachineID int64  `mapstructure:"machine_id"`
	}

	// DurationConfig 过期时间相关的配置
	DurationConfig struct {
		Token int64 `mapstructure:"token"`
	}

	// LogConfig 日志配置
	LogConfig struct {
		// SavePath 日志保存地址
		SavePath string
		// FileName 日志文件名称
		FileName string
		// FileExt 日志文件扩展名
		FileExt string
		// MaxSize 日志切割文件的最大大小
		MaxSize int
		// MaxBackUps 保留旧文件的最大个数
		MaxBackUps int
		// MaxAges 保留旧文件的最大天数
		MaxAge int
	}

	// MySQLConfig MySQL数据库配置
	MySQLConfig struct {
		Host              string `mapstructure:"host"`
		User              string `mapstructure:"user"`
		Password          string `mapstructure:"password"`
		DB                string `mapstructure:"db"`
		Port              int    `mapstructure:"port"`
		MaxOpenConn       int    `mapstructure:"max_open_connection"`
		MaxIdleConn       int    `mapstructure:"max_idle_connection"`
		DefaultStringSize uint   `mapstructure:"default_string_size"`
	}

	// EtcdConfig etcd配置
	EtcdConfig struct {
		// address 服务地址
		Address string
	}

	// ServerConfig 配置服务
	ServerConfig struct {
		// Name 服务名
		Name string
		// address 服务地址
		Address string
	}
)

// NewConfig 创建配置实例
func NewConfig() *Config {
	vp := viper.New()
	vp.SetConfigFile("./config/config.yaml")
	//viper.SetConfigName("config") // 1. 设置配置文件名字
	//viper.SetConfigType("yaml")   // 2. 设置文件类型
	vp.AddConfigPath(".") // 3. 配置文件路径
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("no such configs file: %v\n", err)
		} else {
			log.Fatalf("read configs error: %v\n", err)
		}
	}
	return &Config{
		vp: vp,
	}
}

// WithLogConfig 初始化日志配置
func (c *Config) WithLogConfig() *Config {
	logConf := LogConfig{}
	err := c.vp.UnmarshalKey("log", &logConf)
	if err != nil {
		log.Fatalf("读取日志配置文件失败:%v\n", err)
	}
	c.LogConfig = &logConf
	return c
}

// WithMySQLConfig 初始化MySQL数据库配置
func (c *Config) WithMySQLConfig() *Config {
	mysqlConf := MySQLConfig{}
	err := c.vp.UnmarshalKey("mysql", &mysqlConf)
	if err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}
	c.MySQLConfig = &mysqlConf
	return c
}

// WithBaseConfig 初始化基础配置
func (c *Config) WithBaseConfig() *Config {
	baseConf := BaseConfig{}
	err := c.vp.UnmarshalKey("base", &baseConf)
	if err != nil {
		log.Fatalf("读取视频服务配置文件失败:%v\n", err)
	}
	c.BaseConfig = &baseConf
	return c
}

// WithEtcdConfig 初始化etcd配置
func (c *Config) WithEtcdConfig() *Config {
	etcdConf := EtcdConfig{}
	err := c.vp.UnmarshalKey("etcd", &etcdConf)
	if err != nil {
		log.Fatalf("读取视频服务配置文件失败:%v\n", err)
	}
	c.EtcdConfig = &etcdConf
	return c
}

// WithServerConfig 初始化服务配置
func (c *Config) WithServerConfig(server string) *Config {
	serverConf := ServerConfig{}
	err := c.vp.UnmarshalKey("server."+server, &serverConf)
	if err != nil {
		log.Fatalf("读取视频服务配置文件失败:%v\n", err)
	}
	c.ServerConfig = &serverConf
	return c
}

// WithDurationConfig 初始化过期时间相关的配置
func (c *Config) WithDurationConfig() *Config {
	durationConf := DurationConfig{}
	err := c.vp.UnmarshalKey("duration", &durationConf)
	if err != nil {
		log.Fatalf("读取过期时间相关配置文件失败:%v\n", err)
	}
	c.DurationConfig = &durationConf
	return c
}

func Init() {
	once.Do(func() {
		conf = NewConfig().WithBaseConfig().WithLogConfig().WithMySQLConfig().WithEtcdConfig().WithDurationConfig()
	})
}

func ConfInstance() *Config {
	return conf
}
