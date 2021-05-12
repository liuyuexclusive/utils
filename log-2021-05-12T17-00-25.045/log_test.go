package log

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestLog(t *testing.T) {
	Init("./test.log", zapcore.DebugLevel)

	Sugar.Panic("this is a test.log")

	t.Log("this is a mark")
}
