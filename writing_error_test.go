package gologops

import (
	"errors"
	"testing"
)

type testingBadWriter struct{}

var errTestingBadWriter = errors.New("life goes on bra!")

func (testingBadWriter) Write(b []byte) (int, error) {
	return 0, errTestingBadWriter
}

func TestWriteErr(t *testing.T) {
	l := NewLoggerWithWriter(testingBadWriter{})
	for lvlWanted := allLevel; lvlWanted < noneLevel; lvlWanted++ {
		for _, msgWanted := range stringsForTesting {
			err := l.LogC(logLine{level: lvlWanted, localCx: contextForTesting, message: msgWanted})
			if err != errTestingBadWriter {
				t.Errorf("writer error: want %#v, got %#v", errTestingBadWriter, err)
			}
		}
	}
}
