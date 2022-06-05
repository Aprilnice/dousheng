package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	conf *Config
	once sync.Once
	mu   sync.Mutex
)

type (
	// Config 配置文件结构体
	Config struct {
		vp *viper.Viper
		*ServerConfig
		*LogConfig
		*MySQLConfig
		*RedisConfig
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

	RedisConfig struct {
		PoolSize int    `mapstructure:"pool_size"`
		Port     int    `mapstructure:"port"`
		DB       int    `mapstructure:"db"`
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password"`
	}

	// EtcdConfig etcd配置
	EtcdConfig struct {
		// address 服务地址
		Address string
	}

	Server struct {
		// Name 服务名
		Name string
		// address 服务地址
		Address string
		// 机器地址
		MachineID int64 `mapstructure:"machine_id"`
	}

	// ServerConfig 配置服务
	// 这里会有多个服务 应该用map存储
	ServerConfig map[string]*Server
)

// Server 从ServerConfig 中获取指定的Server
func (s *ServerConfig) Server(server string) *Server {
	srv, ok := (*s)[server]
	if !ok {
		log.Fatalf("%s服务配置未初始化", server)
	}
	return srv
}

// put 往ServerConfig添加服务配置信息
// 加锁操作
func (s *ServerConfig) put(server string, srvInstance *Server) bool {
	mu.Lock()
	defer mu.Unlock()
	return func() bool {
		// 已经存在同名服务了 不允许修改
		if _, ok := (*s)[server]; ok {
			return false
		}
		(*s)[server] = srvInstance
		return true
	}()
}

// NewConfig 创建配置实例
func NewConfig(path string) *Config {
	vp := viper.New()
	yaml := fmt.Sprintf("%s/%s", path, "config.yaml")
	vp.SetConfigFile(yaml)
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
	serverConf := make(ServerConfig, 0)
	return &Config{
		vp:           vp,
		ServerConfig: &serverConf,
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
		log.Fatalf("读取Mysql数据库配置文件失败:%v\n", err)
	}
	c.MySQLConfig = &mysqlConf
	return c
}

// WithRedisConfig 初始化MySQL数据库配置
func (c *Config) WithRedisConfig() *Config {
	redisConf := RedisConfig{}
	err := c.vp.UnmarshalKey("redis", &redisConf)
	if err != nil {
		log.Fatalf("读取Redis数据库配置文件失败:%v\n", err)
	}
	c.RedisConfig = &redisConf
	return c
}

// WithBaseConfig 初始化基础配置
func (c *Config) WithBaseConfig() *Config {
	baseConf := BaseConfig{}
	err := c.vp.UnmarshalKey("base", &baseConf)
	if err != nil {
		log.Fatalf("读取基础配置文件失败:%v\n", err)
	}
	c.BaseConfig = &baseConf
	return c
}

// WithEtcdConfig 初始化etcd配置
func (c *Config) WithEtcdConfig() *Config {
	etcdConf := EtcdConfig{}
	err := c.vp.UnmarshalKey("etcd", &etcdConf)
	if err != nil {
		log.Fatalf("读取ETCD服务配置文件失败:%v\n", err)
	}
	c.EtcdConfig = &etcdConf
	return c
}

// WithServerConfig 初始化服务配置
func (c *Config) WithServerConfig(server string) *Config {
	srvInstance := Server{}
	err := c.vp.UnmarshalKey("server."+server, &srvInstance)
	if err != nil {
		log.Fatalf("读取服务配置文件失败:%v\n", err)
	}
	// 往server map 中添加服务实体
	c.ServerConfig.put(server, &srvInstance)
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

func Init(path string) {
	once.Do(func() {
		conf = NewConfig(path).WithBaseConfig().WithLogConfig().
			WithMySQLConfig().WithRedisConfig().
			WithEtcdConfig().WithDurationConfig()
	})
}

func Instance() *Config {
	return conf
}
