package gologops

import (
	"errors"
	"testing"
)

type testingNestedError struct {
	Text  string
	Cause *testingNestedError
}

func (ce testingNestedError) Error() string { return "a nested error instance" }

func TestFormatError(t *testing.T) {
	ce := testingNestedError{"1", &testingNestedError{"2", &testingNestedError{"3", nil}}}
	l := NewLogger()
	testFormatJSON(t, l, logLine{level: ErrorLevel, err: ce, message: "a nested error"}, "a nested error")
}

type testingNotJSONableNError struct{}

func (testingNotJSONableNError) Error() string { return "a not JSONable error" }
func (testingNotJSONableNError) MarshalJSON() ([]byte, error) {
	return nil, errors.New("JSON not supported")
}

func TestFormatNotJSONError(t *testing.T) {
	ce := testingNotJSONableNError{}
	l := NewLogger()
	msg := "error not should be an object"
	testFormatJSON(t, l, logLine{level: ErrorLevel, err: ce, message: msg}, msg)
}
