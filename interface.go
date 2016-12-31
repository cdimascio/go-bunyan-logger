package logger

type Interface interface {
	WithFields(fields Fields) *Entry
	WithField(key string, value interface{}) *Entry
	WithError(err error) *Entry
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Fatal(msg string)
	Error(msg string)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
	Trace(msg string)
}