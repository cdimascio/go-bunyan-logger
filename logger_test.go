package logger

import (
	"testing"
)

func TestInfoWithData(t *testing.T) {
	log := NewLogger("app")
	log.WithFields(Fields{
		"fieldOne": "Value 1",
	}).Infof("%s", "Info message!")
}

func TestWarnWithGlobalFields(t *testing.T) {
	log := NewLogger("app").WithGlobalFields(Fields{
		"myGlobalFieldOne": "Value One",
		"myGlobalFieldTwo": "Value Two",
	})

	log.Warnf("Msg %d %s", 1, "Warn message!")
}

func TestErrorWithGlobalFieldsAndFields(t *testing.T) {
	log := NewLogger("app").WithGlobalFields(Fields{
		"myGlobalFieldOne": "Value One",
		"myGlobalFieldTwo": "Value Two",
	})

	log.WithFields(Fields{
		"fieldOne": "Value 1",
	}).Errorf("Msg %s", "Error message!")
}


func TestMultipleEntriesWithGlobalFields(t *testing.T) {
	log := NewLogger("example-app").WithGlobalFields(Fields{
		"globalFieldOne": "value 1",
		"globalFieldTwo": 2.0,
		"time": "Woah!",
	})

	log.Error("An error message")

	log.Errorf("An error message with status %d %s", 500, "Internal Error")

	log.WithFields(Fields{
	"fieldOne": "value 1",
	}).Errorf("An error message with %d fields, %d of which are global", 3, 2)

}

func TestMultipleEntries(t *testing.T) {
	log := NewLogger("myapp")
	log.Info("An information message")

	log.Infof("An information message from %s", "go-logger!")

	log.WithFields(Fields{
		"fieldOne": "value 1",
	}).Info("An information message with fields")
}

func TestSetLevel(t *testing.T) {
	log := NewLogger("myapp")
	log.Info("An information message")

	log.SetLevel(LevelError)
	log.Info("An information message")
	log.Error("An Error message")

	log.SetLevel(LevelTrace)
	log.Trace("An information message")
	log.Debug("An information message")
	log.Info("An information message")
	log.Warn("An information message")
	log.Error("An Error message")

	log.SetLevel(LevelWarn)
	log.Trace("An information message")
	log.Debug("An information message")
	log.Info("An information message")
	log.Warn("An information message")
	log.Error("An Error message")
}
