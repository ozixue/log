package log

import (
	"fmt"
	"runtime"
	"time"
)

type logLevel uint8

var (
	colTrace1 = [3]uint{1, 42, 30}
	colTrace2 = [3]uint{0, 0, 32}

	colDebug1 = [3]uint{1, 0, 32}
	colDebug2 = [3]uint{0, 0, 32}

	colInfo1 = [3]uint{1, 0, 34}
	colInfo2 = [3]uint{0, 0, 34}

	colWarning1 = [3]uint{1, 0, 33}
	colWarning2 = [3]uint{0, 0, 33}

	colError1 = [3]uint{1, 0, 31}
	colError2 = [3]uint{0, 0, 31}

	colFatal1 = [3]uint{1, 41, 30}
	colFatal2 = [3]uint{0, 0, 31}

	colFilePath = [3]uint{0, 0, 36}
)

const (
	// TRACE trace level 0
	TRACE logLevel = iota
	// DEBUG debug level 1
	DEBUG
	// INFO info level 2
	INFO
	// WARNING warning level 3
	WARNING
	// ERROR error level 4
	ERROR
	// FATAL fatal level 5
	FATAL
)

// Config log config
type Config struct {
	DateLayout string
	Level      logLevel
}

func (conf Config) enable(level logLevel) bool {
	return level >= conf.Level
}

func (conf Config) log(color1, color2, color3 [3]uint, level, format string, a ...interface{}) {
	now := time.Now().Format(conf.DateLayout)
	_, filePath, line := conf.getFileInfo(3)
	fmt.Printf(
		"%c[%d;%d;%dm%s%s%c[0m %c[%d;%d;%dm%s%s%c[0m %c[%d;%d;%dm%s%s:%d%c[0m Message: %s\n",
		0x1B, color1[0], color1[1], color1[2], "", level, 0x1B, // level color
		0x1B, color2[0], color2[1], color2[2], "", now, 0x1B, // date color
		0x1B, color3[0], color3[1], color3[2], "", filePath, line, 0x1B, //runtime file path color
		fmt.Sprintf(format, a...))
}

func (conf Config) getFileInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, fileName, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	// 获取函数名
	funcName = runtime.FuncForPC(pc).Name()
	//fileName = path.Base(file)
	return
}

var Conf Config

// Trace trace
func Trace(format string, a ...interface{}) {
	if Conf.enable(TRACE) {
		Conf.log(colTrace1, colTrace2, colFilePath, "[T]", format, a...)
	}
}

// Debug debug
func Debug(format string, a ...interface{}) {
	if Conf.enable(DEBUG) {
		Conf.log(colDebug1, colDebug2, colFilePath, "[D]", format, a...)
	}
}

// Info info
func Info(format string, a ...interface{}) {
	if Conf.enable(INFO) {
		Conf.log(colInfo1, colInfo2, colFilePath, "[I]", format, a...)
	}
}

// Warning warning
func Warning(format string, a ...interface{}) {
	if Conf.enable(WARNING) {
		Conf.log(colWarning1, colWarning2, colFilePath, "[W]", format, a...)
	}
}

// Error error
func Error(format string, a ...interface{}) {
	if Conf.enable(ERROR) {
		Conf.log(colError1, colError2, colFilePath, "[E]", format, a...)
	}
}

// Fatal fatal
func Fatal(format string, a ...interface{}) {
	if Conf.enable(FATAL) {
		Conf.log(colFatal1, colFatal2, colFilePath, "[F]", format, a...)
	}
}
