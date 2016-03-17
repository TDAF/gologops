package gologops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func getString(m map[string]interface{}, name string) (string, error) {
	obj, ok := m[name]
	if !ok {
		return "", fmt.Errorf("missing field %q", name)
	}
	str, ok := obj.(string)
	if !ok {
		return "", fmt.Errorf("field %q is not an string", name)
	}
	return str, nil
}

// compareError compares the err object used in log ,o1, with
// the respresentation of that object as map[string]interface{}, o2
// If the original error was not 'JSONable', the representation should
// be a string cointaining the string returned by Error() function of o1
func compareError(o1, o2 interface{}) (bool, error) {
	b, err := json.Marshal(o1)
	if err != nil { // not jsonable object? o2 should be a string
		s2, isString := o2.(string)
		if !isString {
			return false, fmt.Errorf("first arg is not jsonable and second is not an string")
		}
		// The original error string should be included
		return strings.Contains(s2, err.Error()), nil
	}
	var o1map map[string]interface{}
	err = json.Unmarshal(b, &o1map)
	if err != nil {
		return false, nil
	}
	return reflect.DeepEqual(o1map, o2), nil
}

func testFormatJSON(t *testing.T, l *Logger, ll logLine, msgWanted string) {
	var (
		obj    map[string]interface{}
		buffer bytes.Buffer
		err    error
	)

	start := time.Now()
	start, err = time.Parse(timeFormat, start.Format(timeFormat))
	if err != nil {
		t.Fatal(err)
	}

	l.format(&buffer, ll)
	res := buffer.Bytes()
	end := time.Now()
	end, err = time.Parse(timeFormat, end.Format(timeFormat))
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(res, &obj)
	if err != nil {
		t.Fatal(err)
	}

	timeStr, err := getString(obj, "time")
	if err != nil {
		t.Error(err)
	}
	timeStamp, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		t.Error(err)
	}
	if timeStamp.Before(start) {
		t.Error("time before start")
	}
	if timeStamp.After(end) {
		t.Error("time after end")
	}

	lvl, err := getString(obj, "lvl")
	if err != nil {
		t.Error(err)
	}
	if lvl != levelNames[ll.level] {
		t.Errorf("level: wanted %q, got %q", ll.level, lvl)
	}

	msg, err := getString(obj, "msg")
	if err != nil {
		t.Error(err)
	}
	if msg != msgWanted {
		t.Errorf("msg: wanted %q, got %q", msgWanted, msg)
	}

	for k, v := range ll.localCx {
		value, err := getString(obj, k)
		if err != nil {
			t.Error(err)
		}
		if v != value {
			t.Errorf("value for local context field %s: wanted %s, got %s", k, v, value)
		}
	}

	var contextFromFunction C
	var function = l.contextFunc.Load().(func() C)
	if function != nil {
		contextFromFunction = function()
	}
	for k, v := range contextFromFunction {
		if _, ok := ll.localCx[k]; !ok {
			value, err := getString(obj, k)
			if err != nil {
				t.Error(err)
			}
			if v != value {
				t.Errorf("value for func context field %s: wanted %s, got %s", k, v, value)
			}
		}
	}

	for k, v := range l.context.Load().(C) {
		if _, ok := ll.localCx[k]; !ok {
			if _, ok := ll.localCx[k]; !ok {
				value, err := getString(obj, k)
				if err != nil {
					t.Error(err)
				}
				if v != value {
					t.Errorf("value for logger context field %s: wanted %s, got %s", k, v, value)
				}
			}
		}
	}

	if ll.err != nil {
		errObj, ok := obj[ErrFieldName]
		if !ok {
			t.Errorf("missing field  %q", ErrFieldName)
		}
		areEqual, err := compareError(ll.err, errObj)
		if err != nil {
			t.Error(err)
		}
		if !areEqual {
			t.Errorf("value for field %s: wanted %s, got %s", ErrFieldName, ll.err, errObj)
		}
	}
}

func TestSimpleMessageJSON(t *testing.T) {
	l := NewLogger()
	for lvlWanted := allLevel; lvlWanted < noneLevel; lvlWanted++ {
		for _, msgWanted := range stringsForTesting {
			testFormatJSON(t, l, logLine{level: lvlWanted, message: msgWanted}, msgWanted)
		}
	}
}

func TestComplexMessage(t *testing.T) {
	l := NewLogger()
	for lvlWanted := allLevel; lvlWanted < noneLevel; lvlWanted++ {
		for i, text := range stringsForTesting {
			format := "%s,%#v,%f"
			params := []interface{}{text, []int{1, i}, float64(i)}
			msgWanted := fmt.Sprintf(format, params...)
			testFormatJSON(t, l,
				logLine{level: lvlWanted, message: format, params: params},
				msgWanted)
		}
	}
}

func TestLocalContext(t *testing.T) {
	l := NewLogger()
	for lvlWanted := allLevel; lvlWanted < noneLevel; lvlWanted++ {
		for _, msgWanted := range stringsForTesting {
			testFormatJSON(t, l,
				logLine{localCx: contextForTesting, level: lvlWanted, message: msgWanted},
				msgWanted)
		}
	}
}

func TestFuncContext(t *testing.T) {
	l := NewLogger()
	l.SetContextFunc(func() C { return contextForTesting })
	for lvlWanted := allLevel; lvlWanted < noneLevel; lvlWanted++ {
		for _, msgWanted := range stringsForTesting {
			testFormatJSON(t, l,
				logLine{level: lvlWanted, message: msgWanted},
				msgWanted)
		}
	}
}

func TestLoggerContext(t *testing.T) {
	l := NewLogger()
	l.SetContext(contextForTesting)
	for lvlWanted := allLevel; lvlWanted < noneLevel; lvlWanted++ {
		for _, msgWanted := range stringsForTesting {
			testFormatJSON(t, l,
				logLine{level: lvlWanted, message: msgWanted},
				msgWanted)
		}
	}
}

func TestAllContexts(t *testing.T) { t.Skip("to be implemented") }

func testLevelC(t *testing.T, levelMethod Level, method contextLogFunc) {
	var buffer bytes.Buffer
	l := NewLoggerWithWriter(&buffer)
	ctx := C{"trying": "something"}
	for loggerLevel := allLevel; loggerLevel <= levelMethod; loggerLevel++ {
		l.SetLevel(loggerLevel)
		buffer.Reset()
		method(l, ctx, "a not very long message")
		if buffer.Len() == 0 {
			t.Errorf("log not written for method %s when level %s", levelNames[levelMethod], levelNames[loggerLevel])
		}
	}
	for loggerLevel := levelMethod + 1; loggerLevel < noneLevel; loggerLevel++ {
		l.SetLevel(loggerLevel)
		buffer.Reset()
		method(l, ctx, "another short message")
		if buffer.Len() > 0 {
			t.Errorf("log written for method %s when level %s", levelNames[levelMethod], levelNames[loggerLevel])
		}
	}
}

func testLevelf(t *testing.T, levelMethod Level, method formatLogFunc) {

	testLevelC(t, levelMethod, func(l *Logger, ctx C, message string, params ...interface{}) {
		method(l, message, params...)
	})
}

func testLevel(t *testing.T, levelMethod Level, method simpleLogFunction) {

	testLevelC(t, levelMethod, func(l *Logger, ctx C, message string, params ...interface{}) {
		method(l, message)
	})
}

func TestInfof(t *testing.T) {
	testLevelf(t, InfoLevel, (*Logger).Infof)
}

func TestInfo(t *testing.T) {
	testLevel(t, InfoLevel, (*Logger).Info)
}

func TestInfoC(t *testing.T) {
	testLevelC(t, InfoLevel, (*Logger).InfoC)
}
