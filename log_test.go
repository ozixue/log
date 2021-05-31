package log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("number: %d", 100)
}

func TestDebug(t *testing.T) {
	Debug("number: %d", 100)
}

func TestTrace(t *testing.T) {
	Trace("number: %d", 100)
}

func TestWarning(t *testing.T) {
	Warning("number: %d", 100)
}

func TestError(t *testing.T) {
	Error("number: %d", 100)
}
func TestFatal(t *testing.T) {
	Fatal("number: %d", 100)
}

func TestConfig(t *testing.T) {

	Conf.DateLayout = "2006/01/02 15:04:05"
	Conf.Level = WARNING
	// conf.Level = TRACE
	Info("number: %d", 100)
	Debug("number: %d", 100)
	Trace("number: %d", 100)
	Warning("number: %d", 100)
	Error("number: %d", 100)
	Fatal("number: %d", 100)
}
