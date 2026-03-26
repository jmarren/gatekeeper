package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/templates"
)

type MinLen struct {
	FieldName string
	FormName  string
	Value     int
	FmtError  string
}

func NewMinLen(f *Field, v *ValidatorSpec) *MinLen {

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

func (m *MinLen) WriteErrVar(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "min_len_err_var", m)
	if err != nil {
		panic(err)
	}
}

func (m *MinLen) WriteErrorInit(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "min_len_err", m)
	if err != nil {
		panic(err)
	}
}

func (m *MinLen) WriteValidation(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "min_len", m)
	if err != nil {
		panic(err)
	}
}
