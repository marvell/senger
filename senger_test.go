package senger

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	log := NewDefaultLogger()
	log.SetLevel(DebugLevel)

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

func TestLevels(t *testing.T) {
	log := NewDefaultLogger()

	tests := []struct {
		in  []string
		out *LoggerLevel
	}{
		{[]string{"DEBUG", "DBG"}, DebugLevel},
		{[]string{"INFO", "INF"}, InfoLevel},
		{[]string{"WARNING", "WARN", "WRN"}, WarnLevel},
		{[]string{"ERROR", "ERR"}, ErrorLevel},
		{[]string{"FATAL", "FAT"}, FatalLevel},
	}

	for _, test := range tests {
		for _, lvl := range test.in {
			out := log.ParseLevel(lvl)
			if out != test.out {
				t.Errorf("ParseLevel(%s) = %s, want: %s", lvl, out, test.out)
			}
		}
	}
}
