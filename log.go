package beauty

import (
	"log"
	"os"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var logger = NewSimpleLog("beauty", InfoLevel)

type SimpleLog struct {
	log   *log.Logger
	Level LogLevel
}

func NewSimpleLog(module string, logLevel LogLevel) *SimpleLog {
	return &SimpleLog{log: log.New(os.Stderr, "["+module+"]", log.LstdFlags), Level: InfoLevel}
}
func (sl *SimpleLog) SetLevel(logLevel LogLevel) {
	sl.Level = logLevel
}
func SetLogLevel(logLevel LogLevel) {
	logger.SetLevel(logLevel)
}
func (log *SimpleLog) Debug(args ...interface{}) {
	if DebugLevel < log.Level {
		return
	}
	log.log.Println(args)
}
func (log *SimpleLog) Error(args ...interface{}) {
	if ErrorLevel < log.Level {
		return
	}
	log.log.Println(args)
}
func (log *SimpleLog) Println(args ...interface{}) {
	log.log.Println(args)
}
func (log *SimpleLog) Warn(args ...interface{}) {
	if WarnLevel < log.Level {
		return
	}
	log.log.Println(args)
}

func (log *SimpleLog) Panic(args ...interface{}) {
	if PanicLevel < log.Level {
		return
	}
	log.log.Panicln(args)
}

func (log *SimpleLog) Fatal(args ...interface{}) {
	if FatalLevel < log.Level {
		return
	}
	log.log.Fatalln(args)
}

func (log *SimpleLog) Info(args ...interface{}) {
	if InfoLevel < log.Level {
		return
	}
	log.log.Println(args)
}
