package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/util"
)

type MinLen struct {
	FieldName string
	FormName  string
	Value     int
	FmtError  string
	w         io.Writer
}

func (m *MinLen) imports() util.StringSet {
	return util.NewStringSet()
}

func NewMinLen(f *FieldSpec, v *ValidatorSpec) *MinLen {

	val, ok := v.Value.(int)

	if !ok {
		panic("minLen value must be an int")
	}

	return &MinLen{
		FieldName: f.Name,
		FormName:  f.FormName,
		Value:     val,
		FmtError:  v.FmtErr,
	}
}
