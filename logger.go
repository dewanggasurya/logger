package logger

import (
	"io"
	"strings"
	"sync"
	"time"
)

// Level type
type Level int

// List of log level
const (
	DebugLevel   Level = 1
	WarningLevel Level = 2
	InfoLevel    Level = 3
	ErrorLevel   Level = 4
	FatalLevel   Level = 5
	PanicLevel   Level = 6
	UnknownLevel Level = -1
)

var (
	mu       sync.Mutex
	levelMap = map[string]Level{}
	label    = map[Level]string{
		DebugLevel:   "DEBUG",
		WarningLevel: "WARNING",
		InfoLevel:    "INFO",
		ErrorLevel:   "ERROR",
		FatalLevel:   "FATAL",
		PanicLevel:   "PANIC",
	}
)

// ParseLevel from string
func ParseLevel(level string) Level {
	mu.Lock()
	defer mu.Unlock()

	if v, ok := levelMap[strings.ToLower(level)]; ok {
		return v
	}

	return UnknownLevel
}

func init() {
	for k, v := range label {
		levelMap[strings.ToLower(v)] = k
	}
}

// Log struct
type Log struct {
	Time       time.Time `json:"time,omitempty"`
	Message    string    `json:"message,omitempty"`
	Level      Level     `json:"level,omitempty"`
	LevelLabel string    `json:"level_label,omitempty"`
	File       string    `json:"file,omitempty"`
	Line       int       `json:"line,omitempty"`
}

// Logger interface
type Logger interface {
	SetPrefix(prefix string)
	// SetLevel set minimum output log level.
	// The following level are based on its rank, from low to high
	// [Debug, Warning, Info, Error, Fatal, Panic]
	SetLevel(Level)
	SetOutput(io.Writer)
	SetTemplate(string) error
	SetTemplateFormatter(fn func(Log) string)
	ParseLevel(level string) Level
	Write(level Level, calldepth int, message string) error
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})
	Warning(...interface{})
	Warningln(...interface{})
	Warningf(string, ...interface{})
	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})
	Panic(...interface{})
	Panicln(...interface{})
	Panicf(string, ...interface{})
}
