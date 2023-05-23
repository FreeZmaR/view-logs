package logger

import "strings"

const (
	LevelUnknown Level = 0
	LevelInfo    Level = 1
	LevelError   Level = 2
	LevelWarning Level = 3
	LevelFatal   Level = 4
)

type Level int

func ParseLevel(rawLevel string) (Level, bool) {
	level := strings.ToLower(rawLevel)

	if strings.Contains(level, "info") {
		return LevelInfo, true
	}

	if strings.Contains(level, "error") || strings.Contains(level, "err") {
		return LevelError, true
	}

	if strings.Contains(level, "warning") || strings.Contains(level, "warn") {
		return LevelWarning, true
	}

	if strings.Contains(level, "fatal") {
		return LevelFatal, true
	}

	return LevelUnknown, false
}
