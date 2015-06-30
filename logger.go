package senger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"

	sentry "github.com/getsentry/raven-go"
)

const DefaultFlags = log.Ldate | log.Ltime | log.Lmicroseconds

type Logger struct {
	l     *log.Logger
	level *LoggerLevel

	sentry      *sentry.Client
	sentryLevel *LoggerLevel
}

var CallerColor = nMagenta

func NewLogger(w io.Writer, prefix string, lvl *LoggerLevel) *Logger {
	return &Logger{
		l:     log.New(w, prefix, DefaultFlags),
		level: lvl,
	}
}

func NewDefaultLogger() *Logger {
	return NewLogger(os.Stdout, "", InfoLevel)
}

func (l *Logger) EnableSentry(dsn string, lvl *LoggerLevel) error {
	cl, err := newSentryClient(dsn)
	if err != nil {
		return err
	}

	l.sentry = cl
	l.sentryLevel = lvl

	return nil
}

func (l *Logger) SetLevel(lvl *LoggerLevel) {
	l.level = lvl
}

func (l *Logger) isLevel(lvl *LoggerLevel) bool {
	return l.level != nil && l.level.lvl <= lvl.lvl
}

func (l *Logger) isSentryLevel(lvl *LoggerLevel) bool {
	return l.sentryLevel != nil && l.sentryLevel.lvl <= lvl.lvl
}

func (l *Logger) debug(msg string, vars ...interface{}) {
	if l.isSentryLevel(DebugLevel) {
		l.send(msg, DebugLevel)
	}

	if l.isLevel(DebugLevel) {
		msg := fmt.Sprintf("%s %s %s", DebugLevel, caller(3), msg)
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
	if l.isSentryLevel(InfoLevel) {
		l.send(msg, InfoLevel)
	}

	if l.isLevel(InfoLevel) {
		l.l.Print(fmt.Sprintf("%s %s", InfoLevel, msg))
	}
}
func (l *Logger) Infof(f string, vars ...interface{}) {
	l.Info(fmt.Sprintf(f, vars...))
}

func (l *Logger) Warn(msg string) {
	if l.isSentryLevel(WarnLevel) {
		l.send(msg, WarnLevel)
	}

	if l.isLevel(WarnLevel) {
		l.l.Print(fmt.Sprintf("%s %s", WarnLevel, msg))
	}
}
func (l *Logger) Warnf(f string, vars ...interface{}) {
	l.Warn(fmt.Sprintf(f, vars...))
}

func (l *Logger) error(msg interface{}) {
	if l.isSentryLevel(ErrorLevel) {
		l.send(msg.(string), ErrorLevel)
	}

	if l.isLevel(ErrorLevel) {
		msg = fmt.Sprintf("%s %s %s", ErrorLevel, caller(3), msg)
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
	l.send(msg.(string), FatalLevel)

	msg = fmt.Sprintf("%s %s %s", FatalLevel, caller(3), msg)
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

func caller(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	file = path.Base(file)

	return withColor(CallerColor, fmt.Sprintf("%s:%d", file, line))
}
