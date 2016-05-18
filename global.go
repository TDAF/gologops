package gologops

import "io"

// Global logger
var defaultLogger = NewLogger()

func DebugC(context C, message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: DebugLevel, localCx: context, message: message, params: params})
}

func Debugf(message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: DebugLevel, message: message, params: params})
}

func Debug(message string) {
	defaultLogger.LogC(logLine{level: DebugLevel, message: message})
}

func InfoC(context C, message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: InfoLevel, localCx: context, message: message, params: params})
}

func Infof(message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: InfoLevel, message: message, params: params})
}

func Info(message string) {
	defaultLogger.LogC(logLine{level: InfoLevel, message: message})
}

func WarnC(context C, message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: WarnLevel, localCx: context, message: message, params: params})
}

func Warnf(message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: WarnLevel, message: message, params: params})
}

func Warn(message string) {
	defaultLogger.LogC(logLine{level: WarnLevel, message: message})
}

func ErrorE(err error, context C, message string, params ...interface{}) {

	defaultLogger.LogC(logLine{err: err, level: ErrorLevel, localCx: context, message: message, params: params})
}

func ErrorC(context C, message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: ErrorLevel, localCx: context, message: message, params: params})
}

func Errorf(message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: ErrorLevel, message: message, params: params})
}

func Error(message string) {
	defaultLogger.LogC(logLine{level: ErrorLevel, message: message})
}

func FatalC(context C, message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: CriticalLevel, localCx: context, message: message, params: params})
}

func Fatalf(message string, params ...interface{}) {
	defaultLogger.LogC(logLine{level: CriticalLevel, message: message, params: params})
}

func Fatal(message string) {
	defaultLogger.LogC(logLine{level: CriticalLevel, message: message})
}

func FatalE(err error, context C, message string, params ...interface{}) {

	defaultLogger.LogC(logLine{err: err, level: CriticalLevel, localCx: context, message: message, params: params})
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
