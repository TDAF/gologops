package gologops

import "io"

// Global logger
var defaultLogger = NewLogger()

func InfoC(context C, message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: InfoLevel, localCx: context, message: message, params: params})
}

func Infof(message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: InfoLevel, message: message, params: params})
}

func Info(message string) {
	defaultLogger.LogC(logLine{level: InfoLevel, message: message})
}

func SetLevel(lvl Level) {
	defaultLogger.SetLevel(lvl)
}

func SetContext(c C) {
	defaultLogger.SetContext(c)
}

func SetContextFunc(f func() C) {
	defaultLogger.SetContextFunc(f)
}

func SetWriter(w io.Writer) {
	defaultLogger.SetWriter(w)
}

func SetFlags(flags int32) {
	defaultLogger.SetFlags(flags)
}
