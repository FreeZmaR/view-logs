package logger

import (
	"encoding/json"
	"time"
)

type Log interface {
	Parse(data []byte) error

	GetData() string
	GetInfoFields() []*Field
	GetFieldByLabel(label string) (*Field, bool)
	GetLevel() Level
	GetMessage() string
	GetTime() time.Time

	IsLevel(level Level) bool
	HasField(label string, value string, compareType CompareType) bool

	SetInfoFields(labels []string)
}

type Field struct {
	Label  string
	Value  any
	isInfo bool
	layer  int
}

type FieldType interface {
	string | json.Number | bool
}
