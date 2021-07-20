package logger

import (
	"testing"
)

var l = Single().Sugar()

func TestZap(t *testing.T) {
	l.DPanic("this is a test fatal error")
	t.Log("hello")
}
