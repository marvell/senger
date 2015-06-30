package senger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

const (
	DBG_LEVEL = 1<<iota - 1
	INF_LEVEL
	WRN_LEVEL
	ERR_LEVEL
	FAT_LEVEL
)

const (
	DBG = "DBG"
	INF = "INF"
	WRN = "WRN"
	ERR = "ERR"
	FAT = "FAT"
)

const DEFAULT_FLAGS = log.Ldate | log.Ltime | log.Lmicroseconds

type Logger struct {
	l *log.Logger

	level int
}

func NewLogger(w io.Writer, prefix string, lvl int) *Logger {
	return &Logger{
		l:     log.New(w, prefix, DEFAULT_FLAGS),
		level: lvl,
	}
}

func NewDefaultLogger() *Logger {
	return NewLogger(os.Stdout, "", INF_LEVEL)
}

func (l *Logger) SetLevel(lvl int) {
	l.level = lvl
}

func (l *Logger) debug(msg string, vars ...interface{}) {
	if l.level <= DBG_LEVEL {
		file, line := caller(3)
		msg := fmt.Sprintf("[%s] %s:%d %s", DBG, file, line, msg)
		for i, v := range vars {
			msg += fmt.Sprintf("\n\t%d: %#v", i, v)
		}

		l.l.Print(msg)
	}
}
func (l *Logger) Debug(msg string, vars ...interface{}) {
	l.debug(msg, vars...)
}
func (l *Logger) Debugf(f string, vars ...interface{}) {
	l.debug(fmt.Sprintf(f, vars...))
}

func (l *Logger) Info(msg string) {
	if l.level <= INF_LEVEL {
		l.l.Print(fmt.Sprintf("[%s] %s", INF, msg))
	}
}
func (l *Logger) Infof(f string, vars ...interface{}) {
	l.Info(fmt.Sprintf(f, vars...))
}

func (l *Logger) Warn(msg string) {
	if l.level <= WRN_LEVEL {
		l.l.Print(fmt.Sprintf("[%s] %s", WRN, msg))
	}
}
func (l *Logger) Warnf(f string, vars ...interface{}) {
	l.Warn(fmt.Sprintf(f, vars...))
}

func (l *Logger) error(msg interface{}) {
	if l.level <= ERR_LEVEL {
		file, line := caller(3)
		msg = fmt.Sprintf("[%s] %s:%d %s", ERR, file, line, msg)
		l.l.Print(msg)
	}
}
func (l *Logger) Error(msg interface{}) {
	l.error(fmt.Sprintf("%s", msg))
}
func (l *Logger) Errorf(f string, vars ...interface{}) {
	l.error(fmt.Sprintf(f, vars...))
}

func (l *Logger) fatal(msg interface{}) {
	file, line := caller(3)
	msg = fmt.Sprintf("[%s] %s:%d %s", FAT, file, line, msg)
	l.l.Print(msg)
}
func (l *Logger) Fatal(msg interface{}) {
	l.fatal(fmt.Sprintf("%s", msg))
	os.Exit(1)
}
func (l *Logger) Fatalf(f string, vars ...interface{}) {
	l.fatal(fmt.Sprintf(f, vars...))
	os.Exit(1)
}

func caller(skip int) (file string, line int) {
	_, file, line, _ = runtime.Caller(skip)
	file = path.Base(file)

	return
}
