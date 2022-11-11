package logger

import (
	"go-hub/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.Logger

// InitLogger 初始化日志
func InitLogger(level, logType, path string, maxSize, maxBackup, maxAge int, compress bool) {
	// 日志存储格式
	enc := getEncoder()

	// 日志记录介质
	ws := getLogWriter(logType, path, maxSize, maxBackup, maxAge, compress)

	// 设置日志等级
	enab := new(zapcore.Level)
	if err := enab.UnmarshalText([]byte(level)); err != nil {
		panic("日志初始化错误，日志级别设置有误")
	}

	// 初始化 core
	core := zapcore.NewCore(enc, ws, enab)

	// 初始化 Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	// 本地开发环境
	if config.Cfg.Application.Mode == "local" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 终端输出的关键词高亮
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 日志记录介质
func getLogWriter(logType, path string, maxSize, maxBackup, maxAge int, compress bool) zapcore.WriteSyncer {
	filename := ""

	// 按日期记录日志文件
	if logType == "daily" {
		logName := time.Now().Format("2006-01-02") + ".log"
		filename = path + "/" + logName
	} else {
		filename = path + "/logs.log"
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	// 本地开发环境
	if config.Cfg.Application.Mode == "local" {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}

	return zapcore.AddSync(lumberJackLogger)
}
