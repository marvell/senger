package senger

import (
	"testing"
)

func TestLevels(t *testing.T) {
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
			out := ParseLevel(lvl)
			if out != test.out {
				t.Errorf("ParseLevel(%s) = %s, want: %s", lvl, out, test.out)
			}
		}
	}
}
