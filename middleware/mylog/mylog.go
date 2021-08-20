/*
 * https://raw.githubusercontent.com/tsuyopon/golang/master/basic_log/sample6.go
 */
package mylog

import (
    "log"
    "os"
)

type DebugLevel int

const (
    DEBUG DebugLevel = iota // == 0
    INFO  = iota // == 1
    WARN  = iota // == 2
    ERROR = iota // == 3
    FATAL = iota // == 4
)

type Logger interface {
    Debugf(string, ...interface{})
    Infof(string, ...interface{})
    Warnf(string, ...interface{})
    Errorf(string, ...interface{})
    Fatalf(string, ...interface{})
    SetLogLevel(c DebugLevel)
    SetLogger(l Logger)
}

var Mylogger Logger = &BasicLogger{
    Logger: log.New(os.Stderr, "", log.LstdFlags),
}

func (bl *BasicLogger) SetLogger(l Logger) {
    Mylogger = l
}

// デフォルトログレベルはINFO(ライブラリ内スコープ変数なので小文字)
var loglevel DebugLevel = INFO

type BasicLogger struct {
    Logger *log.Logger
}

func (bl *BasicLogger) SetLogLevel(c DebugLevel) {
    loglevel = c
}

func (bl *BasicLogger) Debugf(format string, v ...interface{}) {
	if loglevel > DEBUG {
		return
	}
    bl.Logger.Printf("[DEBUG] " + format, v...)
}

func (bl *BasicLogger) Infof(format string, v ...interface{}) {
	if loglevel > INFO {
		return
	}
    bl.Logger.Printf("[INFO] " + format, v...)
}

func (bl *BasicLogger) Warnf(format string, v ...interface{}) {
	if loglevel > WARN {
		return
	}
    bl.Logger.Printf("[WARN] " + format, v...)
}

func (bl *BasicLogger) Errorf(format string, v ...interface{}) {
	if loglevel > ERROR {
		return
	}
    bl.Logger.Printf("[ERROR] " + format, v...)
}

func (bl *BasicLogger) Fatalf(format string, v ...interface{}) {
	if loglevel > FATAL {
		return
	}
    bl.Logger.Printf("[FATAL] " + format, v...)
}

