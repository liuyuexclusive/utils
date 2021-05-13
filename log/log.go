package log

import (
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/yuexclusive/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// // LogToES logrus to elastic
// func LogrusToES(index string) error {
// 	client, err := es.Client()
// 	if err != nil {
// 		logrus.Error(err)
// 		return err
// 	}
// 	hook, err := elogrus.NewElasticHook(client, config.MustGet().ElasticURL, logrus.InfoLevel, index)
// 	if err != nil {
// 		return fmt.Errorf("fail of log to elastic : %w", err)
// 	}
// 	logrus.AddHook(hook)
// 	gin.DefaultWriter = logrus.StandardLogger().Writer()
// 	logrus.Info("seccessfully of log to elastic")
// 	return nil
// }

var Sugar *zap.SugaredLogger

var Logger *zap.Logger

type level zapcore.Level

func (l level) Enabled(lvl zapcore.Level) bool {
	return l == level(lvl)
}

func Init() {
	encoder := getEncoder()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, getLogWriter(zapcore.DebugLevel.String()), level(zapcore.DebugLevel)),
		zapcore.NewCore(encoder, getLogWriter(zapcore.InfoLevel.String()), level(zapcore.InfoLevel)),
		zapcore.NewCore(encoder, getLogWriter(zapcore.WarnLevel.String()), level(zapcore.WarnLevel)),
		zapcore.NewCore(encoder, getLogWriter(zapcore.ErrorLevel.String()), level(zapcore.ErrorLevel)),
		zapcore.NewCore(encoder, getLogWriter(zapcore.DPanicLevel.String()), level(zapcore.DPanicLevel)),
		zapcore.NewCore(encoder, getLogWriter(zapcore.PanicLevel.String()), level(zapcore.PanicLevel)),
		zapcore.NewCore(encoder, getLogWriter(zapcore.FatalLevel.String()), level(zapcore.FatalLevel)),
	)

	logger := zap.New(core, zap.AddCaller())

	Logger = logger
	Sugar = logger.Sugar()
}

func getLogWriter(level string) zapcore.WriteSyncer {
	hook, err := rotatelogs.New(
		path.Join(config.MustGet().LogPath, "%Y-%m-%d-%H_"+fmt.Sprintf("%s.log", level)),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(time.Hour*24*7),
	)

	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(hook)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
