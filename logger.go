package pluto_go_logger

import (
	"fmt"
	syslog "log"
	"time"
)

// 日志结构体
type Log struct {
	code  uint
	msg   string
	level uint8
	data  []interface{}
}

var showLogLevel = ALL

const (
	ALL uint8 = iota
	TRACE
	DEBUG
	INFO
	WARNING
	FATAL
)

func create(level uint8, msg string, code uint, data ...interface{}) *Log {
	log := new(Log)
	log.level = level
	log.msg = msg
	log.code = code
	log.data = data
	return log
}

// 所有日志
func All(msg string, code uint, data ...interface{}) {
	log := create(DEBUG, msg, code, data...)
	log.handle()
}

// 跟踪日志
func Trace(msg string, code uint, data ...interface{}) {
	log := create(TRACE, msg, code, data...)
	log.handle()
}

// 调试日志
func Debug(msg string, code uint, data ...interface{}) {
	log := create(DEBUG, msg, code, data...)
	log.handle()
}

// 信息日志
func Info(msg string, code uint, data ...interface{}) {
	log := create(INFO, msg, code, data...)
	log.handle()
}

// 警告日志
func Warning(msg string, code uint, data ...interface{}) {
	log := create(WARNING, msg, code, data...)
	log.handle()
}

// 致命日志
func Fatal(msg string, code uint, data ...interface{}) {
	log := create(FATAL, msg, code, data...)
	log.handle()
}

func (log *Log) handle() {

	if log.level < showLogLevel {
		return
	}
	if log.data != nil {
		fmt.Printf("%s [%s]\t[%d] %s %#v \n", time.Now().Format("2006/01/02 15:04:05"), GetLabel(log.level), log.code, log.msg, log.data)
	} else {
		fmt.Printf("%s [%s]\t[%d] %s \n", time.Now().Format("2006/01/02 15:04:05"), GetLabel(log.level), log.code, log.msg)
	}

	if log.level >= FATAL {
		syslog.Fatal(log.code, log.level, log.msg)
	}

}

func GetLabel(levelType uint8) string {

	level := ""
	switch levelType {
	case ALL:
		level = "ALL"
	case TRACE:
		level = "TRACE"
	case DEBUG:
		level = "DEBUG"
	case INFO:
		level = "INFO"
	case WARNING:
		level = "WARNING"
	case FATAL:
		level = "FATAL"
	}
	return level
}

func SetLevel(l uint8) {
	showLogLevel = l
}
func GetLevel() uint8 {
	return showLogLevel
}
