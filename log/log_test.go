package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Init()

	Sugar.Debug("this is a debug log")

	Sugar.Info("this is a info log")

	// t.Log("this is a mark")
}
