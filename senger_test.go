package senger

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	log := NewDefaultLogger()
	log.SetLevel(DBG_LEVEL)

	log.Debug("debug")
	log.Debug("debug with vars", 1, 2, 3)
	log.Debugf("%#v", "debugf")

	log.Info("info")
	log.Infof("%#v", "infof")

	log.Warn("warn")
	log.Warnf("%#v", "warnf")

	log.Error("error string")
	log.Errorf("%#v", "errorf")
	log.Error(errors.New("error error"))

	// log.Fatal("fatal string")
	// log.Fatalf("%#v", "fatalf")
	// log.Fatal(errors.New("fatal error"))
}
