
##go-bunyan-logger
**go-bunyan-logger** is a simple and fast JSON structured-logging library for Go. It is inspired by [bunyan](https://github.com/trentm/node-bunyan).

```go
	log := logger.NewLogger("myapp")
	log.Info("An information message")
```

**go-bunyan-logger** seconds the bunyan manifesto: erver logs should be structured. JSON's a good format. Let's do that. A log record is one line of JSON.stringify'd output. Let's also specify some common names for the requisite and common fields for a log record (see below).

### Compatible with [bunyan's CLI](https://github.com/trentm/node-bunyan)

![](https://github.com/cdimascio/go-bunyan-logger/raw/master/withbunyan.png)

Sample code:

```go
package main

import "github.com/cdimascio/go-bunyan-logger"

var l = logger.NewLogger("my-app").SetLevel(logger.LevelTrace)

func main() {
	l.WithFields(logger.Fields{
		"key": "value",
		"key2": 10,
	}).Info("Woop!")


	l.Errorf("Error: %d (%s)", 500, "Internal Server Error")
	l.Fatal("Fatal message")
	l.Warn("Warning message")
	l.Debug("Debug message")
	l.Trace("Trace message")
}
```



## Installation
`go get github.com/cdimascio/go-bunyan-logger`

## Usage

Quick start example:

```go
package main

import (
  "github.com/cdimascio/go-bunyan-logger"
)

func main() {
	// Create a logger instance. Note: you can create many instances
	log := logger.NewLogger("myapp")

	// Log messages
	log.Info("An information message")

	// Log messages in Printf style
	log.Infof("An information message from %s", "go-bunyan-logger!")

	// Log messages with custom fields
	log.WithFields(logger.Fields{
		"fieldOne": "value 1",
	}).Infof("Msg %s", "Info message!")

}
```

Output:

```json
{"hostname":"Carmines-MacBook-Pro-7.local","level":30,"msg":"An information message","name":"myapp","pid":38851,"time":"2016-12-31T11:34:52-05:00","v":0}

{"hostname":"Carmines-MacBook-Pro-7.local","level":30,"msg":"An information message from [go-bunyan-logger!]","name":"myapp","pid":38851,"time":"2016-12-31T11:34:52-05:00","v":0}

{"fieldOne":"value 1","hostname":"Carmines-MacBook-Pro-7.local","level":30,"msg":"An information message with fields","name":"myapp","pid":38851,"time":"2016-12-31T11:34:52-05:00","v":0}


```

**With Global fields**


```go
   // apply global field to all log entries for this logger instance
  	log := logger.NewLogger().WithGlobalFields(Fields{
		"globalFieldOne": "value 1",
		"globalFieldTwo": 2.0,
	})

	log.Error("An error message")

	log.Errorf("An error message with status %d", 500)

	log.WithFields(Fields{
		"fieldOne": "value 1",
	}).Errorf("An error message with %d fields, %d of which are global", 3, 2)
```

Output:

```json
{"globalFieldOne":"value 1","globalFieldTwo":2,"hostname":"Carmines-MacBook-Pro-7.local","level":50,"msg":"An error message","name":"example-app","pid":38851,"time":"2016-12-31T11:34:52-05:00","v":0}

{"globalFieldOne":"value 1","globalFieldTwo":2,"hostname":"Carmines-MacBook-Pro-7.local","level":50,"msg":"An error message with status [500]","name":"example-app","pid":38851,"time":"2016-12-31T11:34:52-05:00","v":0}

{"fieldOne":"value 1","globalFieldOne":"value 1","globalFieldTwo":2,"hostname":"Carmines-MacBook-Pro-7.local","level":50,"msg":"An error message with 3 fields, 2 of which are global","name":"example-app","pid":38851,"time":"2016-12-31T11:34:52-05:00","v":0}


```

##Configure
The Logger's level and timeformat are configurable

```go
	// Set the logger level. LevelInfo is the default
	log.SetLevel(logger.LevelInfo)

	// defaults to Bunyan time
	log.SetTimeFormat(time.RFC3339)

```

##Methods
```go
	Fatal(msg string)
	Error(msg string)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
	Trace(msg string)

	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Tracef(format string, args ...interface{})

	WithFields(fields Fields) *Entry
	WithField(key string, value interface{}) *Entry
	WithError(err error) * Entry
```


##Levels

The log levels in go-bunyan-logger are as follows. The level descriptions are best practice opinions of the [bunyan](https://github.com/trentm/node-bunyan) author.

- "fatal" (60): The service/app is going to stop or become unusable now. An operator should definitely look into this soon.
- "error" (50): Fatal for a particular request, but the service/app continues servicing other requests. An operator should look at this soon(ish).
- "warn" (40): A note on something that should probably be looked at by an operator eventually.
- "info" (30): Detail on regular operation.
- "debug" (20): Anything else, i.e. too verbose to be included in "info" level.
- "trace" (10): Logging from external libraries used by your app or very detailed application logging.

Change the Logger instance level:

```
	log.SetLevel(logger.LevelInfo)
```

## Core fields
Core fields are provided in each and every log entry.

| field    | required | type   | can be overriden | description                                                                                             | default value |
|----------|----------|--------|------------------|---------------------------------------------------------------------------------------------------------|---------------|
| `v`        | yes      | number    | no               | The go-bunyan-logger format version. The current value is 0.                                                   | `0`             |
| `name`     | yes      | string | yes              | The Logger name. Provided at Logger creation. You must specify a name for your logger when creating it. | n/a           |
| `hostname` | yes      | string | yes              | The hostname. Determined at Logger creation and populated automatically.                                | `os.Hostname()` |
| `pid`      | yes      | number    | no               | The process id. Populated automatically.                                                                | `os.Getpid()`   |
| `time`     | yes      | string | yes              | The timestamp. Populated automatically                                                                  | `time.Now()`    |
| `msg`      | yes      | string | yes              | Every log method e.g. `log.Debug(...)` must provide a msg.                                              | n/a           |

Go ahead and add more fields. Nested ones are fine too!

## Contributing
Pull requests welcomed :)

## License
MIT