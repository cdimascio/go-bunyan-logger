package logger

type Interface interface {
	WithFields(fields Fields) *entry
	WithField(key string, value interface{}) *entry
	WithError(err error) *entry
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