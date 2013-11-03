package tests

import (
	"testing"
	"github.com/ProtoML/ProtoML/logger"
	"log"
	"os"
)

type TestLogger struct {
	t *testing.T
}

func (tl *TestLogger) Fatal(args ...interface{}) {
	tl.t.Fatal(args ...)
}

func (tl *TestLogger) Fatalf(format string, args ...interface{}) {
	tl.t.Fatalf(format, args...)
}

func (tl *TestLogger) Print(args ...interface{}) {
	tl.t.Log(args...)
}

func (tl *TestLogger) Printf(format string, args ...interface{}) {
	tl.t.Logf(format, args...)
}

func (tl *TestLogger) SetPrefix(string) {
	return
}

func SetupLogger(t *testing.T) {
	logger.Debug = true
	//logger.Logger = &TestLogger{t}
	logger.Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}
