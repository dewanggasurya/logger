package log

import (
	"io"

	"github.com/dewanggasurya/logger"
)

// SetPrefix func
func SetPrefix(prefix string) logger.Logger {
	return logger.SetPrefix(prefix)
}

// SetLevel func
func SetLevel(level logger.Level) logger.Logger {
	return logger.SetLevel(level)
}

// SetOutput func
func SetOutput(w io.Writer) logger.Logger {
	return logger.SetOutput(w)
}

// SetTemplate func
func SetTemplate(template logger.Template) logger.Logger {
	return logger.SetTemplate(template)
}

// SetTemplateFormatter func
func SetTemplateFormatter(fn func(logger.Log) string) logger.Logger {
	return logger.SetTemplateFormatter(fn)
}

// Write func
func Write(level logger.Level, callDepth int, message string) error {
	return logger.Write(level, callDepth, message)
}

// Debug func
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Debugln func
func Debugln(v ...interface{}) {
	logger.Debugln(v...)
}

// Debugf func
func Debugf(f string, v ...interface{}) {
	logger.Debugf(f, v...)
}

// Warning func
func Warning(v ...interface{}) {
	logger.Warning(v...)
}

// Warningln func
func Warningln(v ...interface{}) {
	logger.Warningln(v...)
}

// Warningf func
func Warningf(f string, v ...interface{}) {
	logger.Warningf(f, v...)
}

// Info func
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infoln func
func Infoln(v ...interface{}) {
	logger.Infoln(v...)
}

// Infof func
func Infof(f string, v ...interface{}) {
	logger.Infof(f, v...)
}

// Error func
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorln func
func Errorln(v ...interface{}) {
	logger.Errorln(v...)
}

// Errorf func
func Errorf(f string, v ...interface{}) {
	logger.Errorf(f, v...)
}

// Fatal func
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalln func
func Fatalln(v ...interface{}) {
	logger.Fatalln(v...)
}

// Fatalf func
func Fatalf(f string, v ...interface{}) {
	logger.Fatalf(f, v...)
}

// Panic func
func Panic(v ...interface{}) {
	logger.Panic(v...)
}

// Panicln func
func Panicln(v ...interface{}) {
	logger.Panicln(v...)
}

// Panicf func
func Panicf(f string, v ...interface{}) {
	logger.Panicf(f, v...)
}
