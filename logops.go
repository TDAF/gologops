package gologops

import (
	"bytes"
	"os"
	"strings"
	"sync"
	"time"
)

type C map[string]string

type Level int

const (
	allLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	CriticalLevel
	noneLevel
)

var levelNames = [...]string{
	allLevel:      "ALL",
	DebugLevel:    "DEBUG",
	InfoLevel:     "INFO",
	WarnLevel:     "WARN",
	ErrorLevel:    "ERROR",
	CriticalLevel: "FATAL",
	noneLevel:     "NONE",
}

const ErrFieldName = "err"

var (
	timeFormat    string
	prefixFormat  string
	fieldFormat   string
	errorFormat   string
	postfixFormat string
)

func init() {
	format := os.Getenv("LOGOPS_FORMAT")
	if strings.ToLower(format) == "dev" {
		setTextFormat()
	} else {
		setJSONFormat()
	}
}

func setJSONFormat() {
	timeFormat = time.RFC3339
	prefixFormat = `{"time":%q, "lvl":%q`
	fieldFormat = ",%q:%q"
	errorFormat = ",%q:%s"
	postfixFormat = `,"msg":%q}`

}

func setTextFormat() {
	timeFormat = "15:04:05.000"
	prefixFormat = "%s %s\t" // time and level
	fieldFormat = " [%s=%s]" // key and value
	errorFormat = " [%s=%s]" // key and value
	postfixFormat = " %s"    // message
}

var bufferPool = sync.Pool{New: func() interface{} { return &bytes.Buffer{} }}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}
func putBuffer(buffer *bytes.Buffer) {
	buffer.Reset()
	bufferPool.Put(buffer)
}

type logLine struct {
	level   Level
	localCx C
	message string
	params  []interface{}
	err     error
}
