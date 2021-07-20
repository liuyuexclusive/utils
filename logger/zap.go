package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/yuexclusive/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level struct {
	zapcore.Level
}

func (l Level) Enabled(inl zapcore.Level) bool {
	return inl == l.Level
}

func newLogger() *zap.Logger {
	encoder := getEncoder()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, getLogWriter(zapcore.DebugLevel.String()), Level{zapcore.DebugLevel}),
		zapcore.NewCore(encoder, getLogWriter(zapcore.InfoLevel.String()), Level{zapcore.InfoLevel}),
		zapcore.NewCore(encoder, getLogWriter(zapcore.WarnLevel.String()), Level{zapcore.WarnLevel}),
		zapcore.NewCore(encoder, getLogWriter(zapcore.ErrorLevel.String()), Level{zapcore.ErrorLevel}),
		zapcore.NewCore(encoder, getLogWriter(zapcore.DPanicLevel.String()), Level{zapcore.DPanicLevel}),
		zapcore.NewCore(encoder, getLogWriter(zapcore.PanicLevel.String()), Level{zapcore.PanicLevel}),
		zapcore.NewCore(encoder, getLogWriter(zapcore.FatalLevel.String()), Level{zapcore.FatalLevel}),
	)

	logger := zap.New(core, zap.AddCaller())

	return logger
}

var _logger *zap.Logger

var once sync.Once

func Single() *zap.Logger {
	once.Do(func() {
		_logger = newLogger()
	})
	return _logger
}

// getLogWriter 写入文件
func getLogWriter(level string) zapcore.WriteSyncer {
	hook, err := rotatelogs.New(
		path.Join(config.MustGet().LogPath, "%Y-%m-%d-%H_"+fmt.Sprintf("%s.log", level)),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(time.Hour*24*7),
	)

	if err != nil {
		log.Fatal(err)
	}
	return zapcore.AddSync(hook)
}

// getEncoder 获取编码格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
