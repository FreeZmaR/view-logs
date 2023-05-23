package logger

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type JSONLog struct {
	level      Level
	message    string
	time       time.Time
	infoFields []*Field
	allFields  []*Field
	data       string
}

var _ Log = (*JSONLog)(nil)

func NewJSONLog() *JSONLog {
	return &JSONLog{}
}

func (log *JSONLog) Parse(data []byte) error {
	var rawFields map[string]any

	if err := json.Unmarshal(data, &rawFields); err != nil {
		return err
	}

	log.allFields = log.parseFields(rawFields, 0)
	log.parseLevel()
	log.parseMessage()
	log.parseTime()
	log.beautifyData(data)

	return nil
}

func (log *JSONLog) GetData() string {
	return log.data
}

func (log *JSONLog) GetInfoFields() []*Field {
	return log.infoFields
}

func (log *JSONLog) GetFieldByLabel(label string) (*Field, bool) {
	for _, field := range log.allFields {
		if field.Label == label {
			return field, true
		}
	}

	return nil, false
}

func (log *JSONLog) GetLevel() Level {
	return log.level
}

func (log *JSONLog) GetMessage() string {
	return log.message
}

func (log *JSONLog) GetTime() time.Time {
	return log.time
}

func (log *JSONLog) IsLevel(level Level) bool {
	return log.level == level
}

func (log *JSONLog) HasField(label string, value string, compareType CompareType) bool {
	for _, field := range log.allFields {
		if field.Label != label {
			continue
		}

		switch compareType {
		case CompareTypeEqual:
			number, ok := field.Value.(json.Number)
			if !ok {
				return field.Value == value
			}

			return number.String() == value
		case CompareTypeNotEqual:
			number, ok := field.Value.(json.Number)
			if !ok {
				return field.Value != value
			}

			return number.String() != value
		case CompareTypeContain:
			switch t := field.Value.(type) {
			case string:
				return strings.Contains(t, value)
			case json.Number:
				return strings.Contains(t.String(), value)
			default:
				return false
			}
		case CompareTypeNotContain:
			switch t := field.Value.(type) {
			case string:
				return !strings.Contains(t, value)
			case json.Number:
				return !strings.Contains(t.String(), value)
			default:
				return false
			}
		case CompareTypeRegexp:
			r, err := regexp.Compile(value)
			if err != nil {
				return false
			}

			switch t := field.Value.(type) {
			case string:
				return r.MatchString(t)
			case json.Number:
				return r.MatchString(t.String())
			default:
				return false
			}
		case CompareTypeGreatThan:
			switch t := field.Value.(type) {
			case string:
				fieldF, err := strconv.ParseFloat(t, 62)
				if nil == err {
					valueN, err := strconv.ParseInt(value, 10, 64)
					if nil == err {
						return fieldF > float64(valueN)
					}

					valueF, err := strconv.ParseFloat(value, 64)
					if nil == err {
						return fieldF > valueF
					}
				}

				return false

			case json.Number:
				filedN, err := t.Int64()
				if nil == err {
					valueN, err := strconv.ParseInt(value, 10, 64)
					if nil == err {
						return filedN > valueN
					}

					valueF, err := strconv.ParseFloat(value, 64)
					if nil == err {
						return float64(filedN) > valueF
					}

					return false
				}

				filedF, err := t.Float64()
				if nil == err {
					valueN, err := strconv.ParseInt(value, 10, 64)
					if nil == err {
						return filedF > float64(valueN)
					}

					valueF, err := strconv.ParseFloat(value, 64)
					if nil == err {
						return filedF > valueF
					}

					return false
				}

				return false
			default:
				return false
			}
		case CompareTypeLessThan:
			switch t := field.Value.(type) {
			case string:
				fieldF, err := strconv.ParseFloat(t, 62)
				if nil == err {
					valueN, err := strconv.ParseInt(value, 10, 64)
					if nil == err {
						return fieldF < float64(valueN)
					}

					valueF, err := strconv.ParseFloat(value, 64)
					if nil == err {
						return fieldF < valueF
					}
				}

				return false

			case json.Number:
				filedN, err := t.Int64()
				if nil == err {
					valueN, err := strconv.ParseInt(value, 10, 64)
					if nil == err {
						return filedN < valueN
					}

					valueF, err := strconv.ParseFloat(value, 64)
					if nil == err {
						return float64(filedN) < valueF
					}

					return false
				}

				filedF, err := t.Float64()
				if nil == err {
					valueN, err := strconv.ParseInt(value, 10, 64)
					if nil == err {
						return filedF < float64(valueN)
					}

					valueF, err := strconv.ParseFloat(value, 64)
					if nil == err {
						return filedF < valueF
					}

					return false
				}

				return false
			default:
				return false
			}
		case CompareTypeBoolean:
			b, ok := field.Value.(bool)
			if !ok {
				return false
			}

			return b == (value == "true")
		}

	}

	return false
}

func (log *JSONLog) SetInfoFields(labels []string) {
	log.infoFields = make([]*Field, 0, len(labels))

	for _, label := range labels {
		for _, field := range log.allFields {
			if field.Label == label {
				field.isInfo = true
				log.infoFields = append(log.infoFields, field)
			}
		}
	}
}

func (log *JSONLog) parseFields(data map[string]any, layer int) []*Field {
	fields := make([]*Field, 0, len(data))

	for label, value := range data {
		number, ok := value.(json.Number)
		if ok {
			fields = append(fields, &Field{
				Label: label,
				Value: number,
				layer: layer,
			})

			continue
		}

		str, ok := value.(string)
		if ok {
			fields = append(fields, &Field{
				Label: label,
				Value: str,
			})

			continue
		}

		arr, ok := value.([]any)
		if ok {
			fields = append(fields, &Field{
				Label: label,
				Value: arr,
			})

			continue
		}

		m, ok := value.(map[string]any)
		if ok {
			fields = append(fields, log.parseFields(m, layer+1)...)

			continue
		}
	}

	return fields
}

func (log *JSONLog) parseLevel() {
	var lvl Level

	for _, field := range log.allFields {
		if field.layer != 0 {
			continue
		}

		str, ok := field.Value.(string)
		if !ok {
			continue
		}

		lvl, ok = ParseLevel(str)
		if ok {
			log.level = lvl

			return
		}
	}
}

func (log *JSONLog) parseMessage() {
	for _, field := range log.allFields {
		if field.layer != 0 {
			continue
		}

		if field.Label != "msg" && field.Label != "message" {
			continue
		}

		str, ok := field.Value.(string)
		if !ok {
			continue
		}

		log.message = str

		return
	}
}

func (log *JSONLog) parseTime() {
	for _, field := range log.allFields {
		if field.layer != 0 {
			continue
		}

		if field.Label != "time" {
			continue
		}

		str, ok := field.Value.(string)
		if !ok {
			continue
		}

		t, err := time.Parse(time.RFC3339, str)
		if err == nil {
			log.time = t

			return
		}
	}
}

func (log *JSONLog) beautifyData(data []byte) {
	log.data = string(data)
}
