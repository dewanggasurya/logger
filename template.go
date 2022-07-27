package logger

import (
	"strings"
)

const (
	messageKey      = "${message}"
	timeKey         = "${time}"
	callerKey       = "${caller}"
	callerShortKey  = "${caller_short}"
	levelKey        = "${level}"
	prefixKey       = "${prefix}"
	defaultTemplate = "${time} ${level} ${message}\n"
	verboseTemplate = "${time} ${level} ${message} [${caller_short}]"
)

var (
	templateKeys = []string{messageKey, timeKey, callerKey, levelKey}
)

// DefaultTemplate for log
func DefaultTemplate() Template {
	template, _ := ParseTemplate(defaultTemplate)
	return template
}

// VerboseTemplate for log
func VerboseTemplate() Template {
	template, _ := ParseTemplate(verboseTemplate)
	return template
}

// Template struct
// TODO execute function adjustment, so it can be output as JSON or maybe put on hub
type Template struct {
	text   string
	keymap map[string]bool
}

// ParseTemplate log from text
func ParseTemplate(text string) (Template, error) {
	t := Template{
		text:   text,
		keymap: map[string]bool{},
	}

	for _, key := range templateKeys {
		t.keymap[key] = strings.Contains(text, key)
	}

	return t, nil
}

// Execute template
func (t Template) Execute(data map[string]string) string {
	text := t.text

	for k, v := range data {
		text = strings.ReplaceAll(text, k, v)
	}

	return text
}

// Has template key
func (t Template) Has(key string) bool {
	_, ok := t.keymap[key]
	return ok
}

// IsEmpty template
func (t Template) IsEmpty() bool {
	return strings.TrimSpace(t.text) == ""
}

func (t Template) String() string {
	return t.text
}
