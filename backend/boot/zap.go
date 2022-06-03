package boot

import (
	"fmt"
	"gift/lg"
	"gift/util"
	"github.com/natefinch/lumberjack"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}

func InitLogger() error {
	if ok, _ := util.PathExists(DefaultConfig.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", DefaultConfig.Zap.Director)
		_ = os.Mkdir(DefaultConfig.Zap.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	now := time.Now().Format("2006-01-02")

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/%s/debug.log", DefaultConfig.Zap.Director, now), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/info.log", DefaultConfig.Zap.Director, now), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/warn.log", DefaultConfig.Zap.Director, now), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/error.log", DefaultConfig.Zap.Director, now), errorPriority),
	}
	DefaultLog = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if DefaultConfig.Zap.ShowLine {
		DefaultLog = DefaultLog.WithOptions(zap.AddCaller())
	}
	Sugar = DefaultLog.Sugar()
	lg.MyLog = DefaultLog.Sugar()
	Sugar.Info("InitLogger ok")
	return nil
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  DefaultConfig.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case DefaultConfig.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case DefaultConfig.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case DefaultConfig.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case DefaultConfig.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if DefaultConfig.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(DefaultConfig.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     180,  // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	if DefaultConfig.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
