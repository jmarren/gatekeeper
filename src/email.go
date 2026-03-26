package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/templates"
)

type Email struct {
	FieldName string
	FormName  string
	FmtError  string
}

func NewEmail(f *Field, v *ValidatorSpec) *Email {

	return &Email{
		FieldName: f.Name,
		FormName:  f.FormName,
		FmtError:  v.FmtErr,
	}
}

func (e *Email) WriteError(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "email_err", e)
	if err != nil {
		panic(err)
	}
}

func (e *Email) WriteValidation(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "email", e)
	if err != nil {
		panic(err)
	}
}
