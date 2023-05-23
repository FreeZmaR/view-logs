package logger

import (
	"encoding/json"
	"testing"
)

func TestJSONLog_HasField(t *testing.T) {
	type testCaseInput struct {
		label       string
		value       string
		compareType CompareType
	}

	type testCase struct {
		name     string
		fields   []*Field
		input    []testCaseInput
		expected bool
	}

	tt := []testCase{
		{
			name: "Success Equal",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "123",
					compareType: CompareTypeEqual,
				},
				{
					label:       "test_1",
					value:       "123",
					compareType: CompareTypeEqual,
				},
				{
					label:       "test_2",
					value:       "123.1",
					compareType: CompareTypeEqual,
				},
			},
			expected: true,
		},
		{
			name: "Failed Equal",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "12",
					compareType: CompareTypeEqual,
				},
				{
					label:       "test_1",
					value:       "12",
					compareType: CompareTypeEqual,
				},
				{
					label:       "test_2",
					value:       "123",
					compareType: CompareTypeEqual,
				},
			},
			expected: false,
		},
		{
			name: "Success Not Equal",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "1234",
					compareType: CompareTypeNotEqual,
				},
				{
					label:       "test_1",
					value:       "1234",
					compareType: CompareTypeNotEqual,
				},
				{
					label:       "test_2",
					value:       "1234.1",
					compareType: CompareTypeNotEqual,
				},
			},
			expected: true,
		},
		{
			name: "Failed Not Equal",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "123",
					compareType: CompareTypeNotEqual,
				},
				{
					label:       "test_1",
					value:       "123",
					compareType: CompareTypeNotEqual,
				},
				{
					label:       "test_3",
					value:       "some",
					compareType: CompareTypeNotEqual,
				},
			},
			expected: false,
		},
		{
			name: "Success Contain",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "some 123 string",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123456"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1123"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "123",
					compareType: CompareTypeContain,
				},
				{
					label:       "test_1",
					value:       "123",
					compareType: CompareTypeContain,
				},
				{
					label:       "test_2",
					value:       "123.1",
					compareType: CompareTypeContain,
				},
			},
			expected: true,
		},
		{
			name: "Failed Contain",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "456",
					compareType: CompareTypeContain,
				},
				{
					label:       "test_1",
					value:       "12456",
					compareType: CompareTypeContain,
				},
				{
					label:       "test_3",
					value:       "123.1",
					compareType: CompareTypeContain,
				},
			},
			expected: false,
		},
		{
			name: "Success Not Contain",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "some 123 string",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1123"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "sme 123 string",
					compareType: CompareTypeNotContain,
				},
				{
					label:       "test_1",
					value:       "34",
					compareType: CompareTypeNotContain,
				},
				{
					label:       "test_2",
					value:       "1.1",
					compareType: CompareTypeNotContain,
				},
			},
			expected: true,
		},
		{
			name: "Failed Not Contain",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "123",
					compareType: CompareTypeNotContain,
				},
				{
					label:       "test_1",
					value:       "23",
					compareType: CompareTypeNotContain,
				},
				{
					label:       "test_3",
					value:       "123",
					compareType: CompareTypeNotContain,
				},
			},
			expected: false,
		},
		{
			name: "Success Regex",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "https://test.com",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  "https://test.com?test=123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_3",
					Value:  json.Number("123.1123"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       `^https:.*$`,
					compareType: CompareTypeRegexp,
				},
				{
					label:       "test_1",
					value:       `^(https|http):\/\/\w*\.com\?test=\d*$`,
					compareType: CompareTypeRegexp,
				},
				{
					label:       "test_2",
					value:       `123`,
					compareType: CompareTypeRegexp,
				},
				{
					label:       "test_3",
					value:       `\d{3}\.\d{4}`,
					compareType: CompareTypeRegexp,
				},
			},
			expected: true,
		},
		{
			name: "Failed Regex",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       `[a-zA-Z]+`,
					compareType: CompareTypeRegexp,
				},
				{
					label:       "test_1",
					value:       `^\d{2}$`,
					compareType: CompareTypeRegexp,
				},
				{
					label:       "test_2",
					value:       `^[1-2]*.1$`,
					compareType: CompareTypeRegexp,
				},
			},
			expected: false,
		},
		{
			name: "Success GreatThan",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  "-123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_3",
					Value:  json.Number("-123.1123"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       `122`,
					compareType: CompareTypeGreatThan,
				},
				{
					label:       "test_1",
					value:       `-124`,
					compareType: CompareTypeGreatThan,
				},
				{
					label:       "test_2",
					value:       "122",
					compareType: CompareTypeGreatThan,
				},
				{
					label:       "test_3",
					value:       `-123.1124`,
					compareType: CompareTypeGreatThan,
				},
			},
			expected: true,
		},
		{
			name: "Failed GreatThan",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "124",
					compareType: CompareTypeGreatThan,
				},
				{
					label:       "test_1",
					value:       "124",
					compareType: CompareTypeGreatThan,
				},
				{
					label:       "test_2",
					value:       "123.2",
					compareType: CompareTypeGreatThan,
				},
			},
			expected: false,
		},
		{
			name: "Success LessThan",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  "-123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_3",
					Value:  json.Number("-123.1123"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "124",
					compareType: CompareTypeLessThan,
				},
				{
					label:       "test_1",
					value:       "-122",
					compareType: CompareTypeLessThan,
				},
				{
					label:       "test_2",
					value:       "124",
					compareType: CompareTypeLessThan,
				},
				{
					label:       "test_3",
					value:       "-123.1122",
					compareType: CompareTypeLessThan,
				},
			},
			expected: true,
		},
		{
			name: "Failed LessThan",
			fields: []*Field{
				{
					Label:  "test",
					Value:  "123",
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  json.Number("123"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "122",
					compareType: CompareTypeLessThan,
				},
				{
					label:       "test_1",
					value:       "122",
					compareType: CompareTypeLessThan,
				},
				{
					label:       "test_2",
					value:       "122.2",
					compareType: CompareTypeLessThan,
				},
			},
			expected: false,
		},
		{
			name: "Success Boolean",
			fields: []*Field{
				{
					Label:  "test",
					Value:  true,
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  false,
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "true",
					compareType: CompareTypeBoolean,
				},
				{
					label:       "test_1",
					value:       "false",
					compareType: CompareTypeBoolean,
				},
			},
			expected: true,
		},
		{
			name: "Failed Boolean",
			fields: []*Field{
				{
					Label:  "test",
					Value:  true,
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_1",
					Value:  false,
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_2",
					Value:  json.Number("123.1"),
					isInfo: false,
					layer:  0,
				},
				{
					Label:  "test_3",
					Value:  "true",
					isInfo: false,
					layer:  0,
				},
			},
			input: []testCaseInput{
				{
					label:       "test",
					value:       "false",
					compareType: CompareTypeBoolean,
				},
				{
					label:       "test_1",
					value:       "true",
					compareType: CompareTypeBoolean,
				},
				{
					label:       "test_2",
					value:       "false",
					compareType: CompareTypeBoolean,
				},
				{
					label:       "test_3",
					value:       "true",
					compareType: CompareTypeBoolean,
				},
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			log := &JSONLog{allFields: tc.fields}

			for _, input := range tc.input {
				if log.HasField(input.label, input.value, input.compareType) != tc.expected {
					t.Errorf(
						"Input Lebel: %s Value: %s CompareType: %s Expected %v, got %v",
						input.label,
						input.value,
						input.compareType.String(),
						tc.expected,
						!tc.expected,
					)
				}
			}

		})
	}
}
