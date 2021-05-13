package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	// 限制日志输出级别, >= DebugLevel 会打印所有级别的日志
	// 生产环境中一般使用 >= ErrorLevel
	lowPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.DebugLevel
	})

	// 使用 JSON 格式日志
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	stdCore := zapcore.NewCore(jsonEnc, zapcore.Lock(os.Stdout), lowPriority)

	// logger 输出到 console 且标识调用代码行
	return zap.New(stdCore).WithOptions(zap.AddCaller())
}

func main() {
	logger := NewLogger()

	logger.Info("test logger info", zap.String("hello", "logger"))
}
