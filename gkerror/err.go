package gkerror

import (
	"fmt"
	"slices"
	"strings"
)

type ValidationErr interface {
	Error() string
	Field() string
	Expected() any
}

type ValidationErrGroup struct {
	errs []ValidationErr
}

func (v *ValidationErrGroup) Print() {
	for _, err := range v.errs {
		fmt.Println(err.Error())
	}
}

func (v *ValidationErrGroup) String() string {
	out := ""
	for _, err := range v.errs {
		out += "\n"
		out += err.Error()
	}
	return out
}

func NewValidationErrGroup() *ValidationErrGroup {
	return new(ValidationErrGroup)
}

func (v *ValidationErrGroup) Errors() []string {
	out := []string{}
	for _, err := range v.errs {
		out = append(out, err.Error())
	}
	return out
}

func (v *ValidationErrGroup) Add(vErr ValidationErr) {
	v.errs = append(v.errs, vErr)
}

func (v *ValidationErrGroup) ByField(field string) *ValidationErrGroup {
	out := NewValidationErrGroup()
	for _, err := range v.errs {
		if err.Field() == field {
			out.Add(err)
		}
	}
	return out
}

// func hi() {
// 	minLenErr := func(received int) ValidationErr {
// 		return NewValidationErr("FirstName", received, 3, "must be >= %(value)")
// 	}
//
// }

type validationErr struct {
	field        string
	received     any
	expected     any
	formatString string
	fmtParams    []any
}

func NewValidationErr(field string, received any, expected any, formatString string) ValidationErr {

	valueIndex := strings.Index(formatString, "%(received)")
	expectedIndex := strings.Index(formatString, "%(expected)")
	fieldIndex := strings.Index(formatString, "%(field)")

	fmtParams := []any{}

	indexes := make(map[int]string)

	for valueIndex != -1 {
		indexes[valueIndex] = "received"
		valueIndex = strings.Index(formatString[valueIndex+1:], "%(received)")
	}

	for expectedIndex != -1 {
		indexes[expectedIndex] = "expected"
		expectedIndex = strings.Index(formatString[expectedIndex+1:], "%(expected)")
	}

	for fieldIndex != -1 {
		indexes[fieldIndex] = "field"
		fieldIndex = strings.Index(formatString[fieldIndex+1:], "%(field)")
	}

	keys := []int{}

	for key := range indexes {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		switch indexes[key] {
		case "received":
			fmtParams = append(fmtParams, received)
		case "expected":
			fmtParams = append(fmtParams, expected)
		case "field":
			fmtParams = append(fmtParams, field)
		}
	}

	formatString = strings.ReplaceAll(formatString, "%(received)", "%v")
	formatString = strings.ReplaceAll(formatString, "%(expected)", "%v")
	formatString = strings.ReplaceAll(formatString, "%(field)", "%s")

	return &validationErr{
		field,
		received,
		expected,
		formatString,
		fmtParams,
	}
}

func (v *validationErr) Error() string {
	fmt.Println(v.formatString)
	return fmt.Sprintf(v.formatString, v.fmtParams...)
}

func (v *validationErr) Field() string {
	return v.field
}

func (v *validationErr) Expected() any {
	return v.expected
}
