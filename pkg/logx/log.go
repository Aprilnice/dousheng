package logx

import (
	"dousheng/config"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

// InitLogger 初始化日志并创建实例
func InitLogger(config config.Config) {
	writeSyncer := getLogWriter(config)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)
	Log = zap.New(core, zap.AddCaller())
}

// getEncoder 获取编码器
func getEncoder() zapcore.Encoder {
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConf)
}

// getLogWriter 获取写同步器
func getLogWriter(conf config.Config) zapcore.WriteSyncer {
	fileName := path.Join(conf.LogConfig.SavePath, conf.LogConfig.FileName+conf.LogConfig.FileExt)

	lumLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    conf.LogConfig.MaxSize,
		MaxBackups: conf.LogConfig.MaxBackUps,
		MaxAge:     conf.LogConfig.MaxAge,
		Compress:   false,
	}

	return zapcore.AddSync(lumLogger)
}
