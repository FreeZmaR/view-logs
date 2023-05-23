package logger

/*
CompareType is an enum for the type of comparison to be done on a field.
Also, CompareType represent json type (string, number, boolean)

CompareTypeEqual: Equal to the user value and the json field value must be string or json.Number
CompareTypeNotEqual: Not equal to the user value and the json field value must be string or json.Number
CompareTypeContain: The json field value must be string or json.Number and contain the user value
CompareTypeNotContain: The json field value must be string or json.Number and not contain the user value
CompareTypeRegexp: The json field value must be string or json.Number and match the user value as a regular expression
CompareTypeGreatThan: The json field value must be json.Number and greater than the user value, user value must be int or float32 by string
CompareTypeLessThan: The json field value must be json.Number and less than the user value, user value must be int or float32 by string
CompareTypeBoolean: The json field value must be boolean and equal to the user value, user value must be (true or false) by string
*/

//go:generate stringer -type=CompareType -output=compare_type_string.go
const (
	CompareTypeEqual CompareType = iota + 1
	CompareTypeNotEqual
	CompareTypeContain
	CompareTypeNotContain
	CompareTypeRegexp
	CompareTypeGreatThan
	CompareTypeLessThan
	CompareTypeBoolean

	CompareEqualText      = "Equal"
	CompareNotEqualText   = "Not Equal"
	CompareContainText    = "Contain"
	CompareNotContainText = "Not Contain"
	CompareRegexpText     = "Regexp"
	CompareGreatThanText  = "Greater Than"
	CompareLessThanText   = "Less Than"
	CompareBooleanText    = "Boolean"
)

type CompareType int

func GetCompareTypeText() []string {
	return []string{
		CompareEqualText,
		CompareNotEqualText,
		CompareContainText,
		CompareNotContainText,
		CompareRegexpText,
		CompareGreatThanText,
		CompareLessThanText,
		CompareBooleanText,
	}
}

func GetCompareTypeByString(str string) CompareType {
	switch str {
	case CompareEqualText:
		return CompareTypeEqual
	case CompareNotEqualText:
		return CompareTypeNotEqual
	case CompareContainText:
		return CompareTypeContain
	case CompareNotContainText:
		return CompareTypeNotContain
	case CompareRegexpText:
		return CompareTypeRegexp
	case CompareGreatThanText:
		return CompareTypeGreatThan
	case CompareLessThanText:
		return CompareTypeLessThan
	case CompareBooleanText:
		return CompareTypeBoolean
	default:
		return CompareTypeEqual
	}
}
