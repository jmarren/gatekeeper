package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/templates"
)

type MaxLen struct {
	FieldName string
	FormName  string
	Value     int
	FmtError  string
}

func NewMaxLen(f *Field, v *ValidatorSpec) *MaxLen {

	val, ok := v.Value.(int)

	if !ok {
		panic("maxLen value must be an int")
	}

	return &MaxLen{
		FieldName: f.Name,
		FormName:  f.FormName,
		Value:     val,
		FmtError:  v.FmtErr,
	}
}

func (m *MaxLen) WriteError(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "max_len_err", m)
	if err != nil {
		panic(err)
	}
}

func (m *MaxLen) WriteValidation(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "max_len", m)
	if err != nil {
		panic(err)
	}
}
