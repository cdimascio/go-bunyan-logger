package logger

import (
	"fmt"
	"os"
	"time"
	"encoding/json"
)


type entry struct {
	logger *Logger
	Fields Fields
}

var systemFields = map[string]interface{}{
	"name":nil,
	"v":nil,
	"level":nil,
	"pid":nil,
}
var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = "na"
		fmt.Fprint(os.Stderr, "Failed to obtain hostname, %v\n", err)
	}
}


func newEntry(logger *Logger) *entry {
	return &entry{
		logger: logger,
		Fields: Fields{
			"name": logger.Name,
			"hostname": hostname,
			"pid": os.Getpid(),
			"time": time.Now().Format(logger.TimeFormat),
			"v": Version,
		},
	}
}

func (e *entry) WithFields(fields Fields) *entry {
	allFields := make(Fields)
	addFields(&allFields, e.logger.Fields)
	addFields(&allFields, e.Fields)
	addFields(&allFields, fields)
	e.Fields = allFields
	return e

}

func addFields(result *Fields, fields Fields) {
	for key, value := range fields {
		if isSystemField(key) {
			continue
		}
		var r = *result
		r[key] = value
	}
}

func (e *entry) WithField(key string, value interface{}) *entry {
	return e.WithFields(Fields{key: value})
}

func (e *entry) WithError(err error) *entry {
	return e.WithField("error", err.Error())
}

func (e *entry) Fatalf(format string, args ...interface{}) {
	e.log(LevelFatal, fmt.Sprintf(format, args...))
}

func (e *entry) Errorf(format string, args ...interface{}) {
	e.log(LevelError, fmt.Sprintf(format, args...))
}

func (e *entry) Warnf(format string, args ...interface{}) {
	e.log(LevelWarn, fmt.Sprintf(format, args...))
}

func (e *entry) Infof(format string, args ...interface{}) {
	e.log(LevelInfo, fmt.Sprintf(format, args...))
}

func (e *entry) Debugf(format string, args ...interface{}) {
	e.log(LevelDebug, fmt.Sprintf(format, args...))
}

func (e *entry) Tracef(format string, args ...interface{}) {
	e.log(LevelTrace, fmt.Sprintf(format, args...))
}

func (e *entry) Fatal(msg string) {
	e.log(LevelFatal, msg)
}

func (e *entry) Error(msg string) {
	e.log(LevelError, msg)
}

func (e *entry) Warn(msg string) {
	e.log(LevelWarn, msg)
}

func (e *entry) Info(msg string) {
	e.log(LevelInfo, msg)
}

func (e *entry) Debug(msg string) {
	e.log(LevelDebug, msg)
}

func (e *entry) Trace(msg string) {
	e.log(LevelTrace, msg)
}

func (e *entry) log(level Level, msg string) {
	if level >= e.logger.Level  {
		entryMap := map[string]interface{}{
			"level":    level,
			"msg":      msg,
		}

		for key, value := range e.Fields {
			entryMap[key] = value
		}

		enc := json.NewEncoder(e.logger.Out)

		e.logger.mutex.Lock()
		defer e.logger.mutex.Unlock()

		enc.Encode(entryMap)
	}
}

func isSystemField(key string) bool {
	_, ok := systemFields[key]
	return ok
}