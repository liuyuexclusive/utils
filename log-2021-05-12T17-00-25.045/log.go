package log

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
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

func Init(logFilePath string, level zapcore.LevelEnabler) {
	writeSyncer := getLogWriter(logFilePath)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, level)

	logger := zap.New(core, zap.AddCaller())
	Logger = logger
	Sugar = logger.Sugar()
}

func getLogWriter(logFilePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func Gin(engine *gin.Engine) {
	engine.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(Logger, true))
}
