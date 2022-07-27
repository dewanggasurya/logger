package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var base Logger

// Base struct
type Base struct {
	mu                  sync.Mutex
	outLevel            Level
	out                 io.Writer
	prefix              string
	template            Template
	templateFormatterFn func(Log) string
	timeFormat          string
	isDiscard           bool
}

func init() {
	base = Default()
}

// SetLogger to package logger
func SetLogger(logger Logger) {
	mu.Lock()
	defer mu.Unlock()

	base = logger
}

// Default logger
func Default() Logger {
	return &Base{
		outLevel:   InfoLevel,
		out:        os.Stderr,
		template:   DefaultTemplate(),
		timeFormat: time.RFC3339,
	}
}

// ParseLevel from string
func (b *Base) ParseLevel(level string) Level {
	return ParseLevel(level)
}

// SetPrefix func
func (b *Base) SetPrefix(prefix string) Logger {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.prefix = prefix
	return b
}

// SetLevel func
func (b *Base) SetLevel(level Level) Logger {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.outLevel = level
	return b
}

// SetOutput func
func (b *Base) SetOutput(w io.Writer) Logger {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.out = w
	b.isDiscard = w == io.Discard
	return b
}

// SetTemplate func
func (b *Base) SetTemplate(template Template) Logger {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.template = template
	return b
}

// SetTemplateFormatter func
func (b *Base) SetTemplateFormatter(fn func(Log) string) Logger {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.templateFormatterFn = fn

	return b
}

// Write func
func (b *Base) Write(level Level, callDepth int, message string) error {
	now := time.Now()

	if level < b.outLevel || b.isDiscard {
		return nil
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.template.IsEmpty() {
		b.template, _ = ParseTemplate(defaultTemplate)
	}

	logData := Log{}
	keyMap := map[string]string{}
	if b.template.Has(timeKey) {
		keyMap[timeKey] = now.Format(b.timeFormat)
		logData.Time = now
	}

	if b.template.Has(callerKey) || b.template.Has(callerShortKey) {
		b.mu.Unlock()
		_, file, line, ok := runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		}
		b.mu.Lock()

		keyMap[callerShortKey] = fmt.Sprintf("%s:%d", filepath.Base(file), line)
		keyMap[callerKey] = fmt.Sprintf("%s:%d", file, line)
		logData.File = file
		logData.Line = line
	}

	if b.template.Has(levelKey) {
		levelLabel := label[level]

		keyMap[levelKey] = fmt.Sprintf("%-10s", levelLabel)
		logData.LevelLabel = levelLabel
		logData.Level = level
	}

	logData.Message = message

	var text string
	if b.templateFormatterFn == nil {
		// putting message replacement outside
		text = b.template.Execute(keyMap)
		text = strings.ReplaceAll(text, messageKey, message)
	} else {
		text = b.templateFormatterFn(logData)
	}

	buf := []byte(text)
	if len(text) > 0 && text[len(text)-1] != '\n' {
		buf = append(buf, '\n')
	}

	_, err := b.out.Write(buf)

	return err
}

// Debug func
func (b *Base) Debug(v ...interface{}) {
	b.Write(DebugLevel, 2, fmt.Sprint(v...))
}

// Debugln func
func (b *Base) Debugln(v ...interface{}) {
	b.Write(DebugLevel, 2, fmt.Sprintln(v...))
}

// Debugf func
func (b *Base) Debugf(f string, v ...interface{}) {
	b.Write(DebugLevel, 2, fmt.Sprintf(f, v...))
}

// Warning func
func (b *Base) Warning(v ...interface{}) {
	b.Write(WarningLevel, 2, fmt.Sprint(v...))

}

// Warningln func
func (b *Base) Warningln(v ...interface{}) {
	b.Write(WarningLevel, 2, fmt.Sprintln(v...))

}

// Warningf func
func (b *Base) Warningf(f string, v ...interface{}) {
	b.Write(WarningLevel, 2, fmt.Sprintf(f, v...))
}

// Info func
func (b *Base) Info(v ...interface{}) {
	b.Write(InfoLevel, 2, fmt.Sprint(v...))
}

// Infoln func
func (b *Base) Infoln(v ...interface{}) {
	b.Write(InfoLevel, 2, fmt.Sprintln(v...))
}

// Infof func
func (b *Base) Infof(f string, v ...interface{}) {
	b.Write(InfoLevel, 2, fmt.Sprintf(f, v...))
}

// Error func
func (b *Base) Error(v ...interface{}) {
	b.Write(ErrorLevel, 2, fmt.Sprint(v...))
}

// Errorln func
func (b *Base) Errorln(v ...interface{}) {
	b.Write(ErrorLevel, 2, fmt.Sprintln(v...))
}

// Errorf func
func (b *Base) Errorf(f string, v ...interface{}) {
	b.Write(ErrorLevel, 2, fmt.Sprintf(f, v...))
}

// Fatal func
func (b *Base) Fatal(v ...interface{}) {
	b.Write(FatalLevel, 2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalln func
func (b *Base) Fatalln(v ...interface{}) {
	b.Write(FatalLevel, 2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf func
func (b *Base) Fatalf(f string, v ...interface{}) {
	b.Write(FatalLevel, 2, fmt.Sprintf(f, v...))
	os.Exit(1)
}

// Panic func
func (b *Base) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	b.Write(PanicLevel, 2, s)
	panic(s)
}

// Panicln func
func (b *Base) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	b.Write(PanicLevel, 2, s)
	panic(s)
}

// Panicf func
func (b *Base) Panicf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	b.Write(PanicLevel, 2, s)
	panic(s)
}

// SetPrefix func
func SetPrefix(prefix string) Logger {
	return base.SetPrefix(prefix)
}

// SetLevel func
func SetLevel(level Level) Logger {
	return base.SetLevel(level)
}

// SetOutput func
func SetOutput(w io.Writer) Logger {
	return base.SetOutput(w)
}

// SetTemplate func
func SetTemplate(template Template) Logger {
	return base.SetTemplate(template)
}

// SetTemplateFormatter func
func SetTemplateFormatter(fn func(Log) string) Logger {
	return base.SetTemplateFormatter(fn)
}

// Write func
func Write(level Level, callDepth int, message string) error {
	return base.Write(level, callDepth, message)
}

// Debug func
func Debug(v ...interface{}) {
	base.Debug(v)
}

// Debugln func
func Debugln(v ...interface{}) {
	base.Debugln(v)
}

// Debugf func
func Debugf(f string, v ...interface{}) {
	base.Debugf(f, v)
}

// Warning func
func Warning(v ...interface{}) {
	base.Warning(v)
}

// Warningln func
func Warningln(v ...interface{}) {
	base.Warningln(v)
}

// Warningf func
func Warningf(f string, v ...interface{}) {
	base.Warningf(f, v)
}

// Info func
func Info(v ...interface{}) {
	base.Info(v)
}

// Infoln func
func Infoln(v ...interface{}) {
	base.Infoln(v)
}

// Infof func
func Infof(f string, v ...interface{}) {
	base.Infof(f, v)
}

// Error func
func Error(v ...interface{}) {
	base.Error(v)
}

// Errorln func
func Errorln(v ...interface{}) {
	base.Errorln(v)
}

// Errorf func
func Errorf(f string, v ...interface{}) {
	base.Errorf(f, v)
}

// Fatal func
func Fatal(v ...interface{}) {
	base.Fatal(v)
}

// Fatalln func
func Fatalln(v ...interface{}) {
	base.Fatalln(v)
}

// Fatalf func
func Fatalf(f string, v ...interface{}) {
	base.Fatalf(f, v)
}

// Panic func
func Panic(v ...interface{}) {
	base.Panic(v)
}

// Panicln func
func Panicln(v ...interface{}) {
	base.Panicln(v)
}

// Panicf func
func Panicf(f string, v ...interface{}) {
	base.Panicf(f, v)
}
