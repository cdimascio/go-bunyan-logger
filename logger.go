package logger

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	out        io.Writer
	name       string
	level      Level
	fields     Fields
	timeFormat string
	mutex      sync.Mutex
}

func NewLogger(name string) *Logger {
	return &Logger{
		out: os.Stdout,
		name: name,
		level: LevelInfo,
		timeFormat: "2006-01-02T15:04:05.000Z",
		fields: nil,
	}
}

func (l *Logger) SetLevel(level Level) *Logger {
	l.level = level
	return l
}

func (l *Logger) WithGlobalFields(fields Fields) *Logger {
	l.fields = fields
	return l
}

func (l *Logger) SetTimeFormat(layout string) *Logger {
	l.timeFormat = layout
	return l
}

func (l *Logger) WithFields(fields Fields) *entry {
	return newEntry(l).WithFields(fields)
}

func (l *Logger) WithField(key string, value interface{}) *entry {
	return newEntry(l).WithField(key, value)
}

func (l *Logger) WithError(err error) *entry {
	return newEntry(l).WithError(err)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	newEntry(l).Fatalf(format,args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	newEntry(l).Errorf(format,args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	newEntry(l).Warnf(format,args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	newEntry(l).Infof(format,args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	newEntry(l).Debugf(format,args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	newEntry(l).Tracef(format,args...)
}

func (l *Logger) Fatal(msg string) {
	newEntry(l).Fatal(msg)
}

func (l *Logger) Error(msg string) {
	newEntry(l).Error(msg)
}

func (l *Logger) Warn(msg string) {
	newEntry(l).Warn(msg)
}

func (l *Logger) Info(msg string) {
	newEntry(l).Info(msg)
}

func (l *Logger) Debug(msg string) {
	newEntry(l).Debug(msg)
}

func (l *Logger) Trace(msg string) {
	newEntry(l).Trace(msg)
}
