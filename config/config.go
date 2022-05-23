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
		*VideoServerConfig
		*LogConfig
		*MySQLConfig
		*BaseConfig
	}

	// BaseConfig 基础配置
	BaseConfig struct {
		Host      string `mapstructure:"host"`
		Name      string `mapstructure:"name"`
		Mode      string `mapstructure:"mode"`
		Version   string `mapstructure:"version"`
		StartTime string `mapstructure:"start_time"`
		Port      int    `mapstructure:"port"`
		MachineID int64  `mapstructure:"machine_id"`
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

	// VideoServerConfig 视频配置服务
	VideoServerConfig struct {
		// Name 服务名
		Name string
		// Port 端口
		Port string
	}
)

// NewConfig 创建配置实例
func NewConfig() *Config {
	vp := viper.New()
	vp.SetConfigFile("config.yaml")
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

// WithVideoConfig 初始化视频服务配置
func (c *Config) WithVideoConfig() *Config {
	videoConf := VideoServerConfig{}
	err := c.vp.UnmarshalKey("server.video", &videoConf)
	if err != nil {
		log.Fatalf("读取视频服务配置文件失败:%v\n", err)
	}
	return &Config{
		VideoServerConfig: &videoConf,
	}
}

func Init() {
	once.Do(func() {
		conf = NewConfig().WithBaseConfig().WithLogConfig().WithMySQLConfig()
	})
}

func ConfInstance() *Config {
	return conf
}
