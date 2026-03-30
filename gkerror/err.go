package gkerror

import (
	"fmt"
	"strings"
)

type ValidationErr interface {
	Error() string
	Field() string
	Expected() any
}

type ValidationErrGroup struct {
	errs   []ValidationErr
	errStr *strings.Builder
}

func (v *ValidationErrGroup) Print() {
	for _, err := range v.errs {
		fmt.Println(err.Error())
	}
}

func (v *ValidationErrGroup) String() string {
	out := new(strings.Builder)
	for _, err := range v.errs {
		out.WriteString("\n" + err.Error())
	}
	return out.String()
}

func NewValidationErrGroup() *ValidationErrGroup {
	return &ValidationErrGroup{
		errs:   []ValidationErr{},
		errStr: new(strings.Builder),
	}
}

func (v *ValidationErrGroup) Any() bool {
	return len(v.errs) > 0
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
	v.errStr.WriteString("\n" + vErr.Error())
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

type validationErr struct {
	field    string
	received any
	expected any
	error    string
}

func NewValidationErr(field string, received any, expected any, formatString string) ValidationErr {
	error := strings.ReplaceAll(formatString, "%v", fmt.Sprintf("%v", received))
	return &validationErr{
		field,
		received,
		expected,
		error,
	}
}

func (v *validationErr) Error() string {
	return v.error
}

func (v *validationErr) Field() string {
	return v.field
}

func (v *validationErr) Expected() any {
	return v.expected
}

type ValidationErrReciever func(received any) ValidationErr

func NewErrReceiver(field string, expected any, formatString string) ValidationErrReciever {
	return func(received any) ValidationErr {
		return NewValidationErr(field, received, expected, formatString)
	}
}
