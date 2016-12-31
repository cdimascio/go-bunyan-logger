package logger

type Level int8

const (
	LevelFatal = 60
	LevelError = 50
	LevelWarn = 40
	LevelInfo = 30
	LevelDebug = 20
	LevelTrace = 10
)

func (level Level) name() string {
	switch level {
	case LevelFatal:
		return "fatal"
	case LevelError:
		return "error"
	case LevelWarn:
		return "warn"
	case LevelInfo:
		return "info"
	case LevelDebug:
		return "debug"
	case LevelTrace:
		return "trace"
	default:
		return "unknown"
	}
}