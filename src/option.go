package src

import (
	"fmt"
	"io"

	"github.com/jmarren/gatekeeper/src/templates"
)

type Option struct {
	FieldName string
	FormName  string
	Value     []string
	FmtError  string
}

func NewOption(f *FieldSpec, v *ValidatorSpec) *Option {

	fmt.Printf("v.Value = %v\n", v.Value)

	vals := []string{}

	iVals, ok := v.Value.([]any)
	if !ok {
		panic("option value must be a list")
	}

	for _, iVal := range iVals {
		val, ok := iVal.(string)
		if !ok {
			panic("option value must be a list of strings")
		}
		vals = append(vals, val)

	}

	return &Option{
		FieldName: f.Name,
		FormName:  f.FormName,
		FmtError:  v.FmtErr,
		Value:     vals,
	}
}

func (o *Option) WriteError(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "option_var", o)
	if err != nil {
		panic(err)
	}

	err = templates.Tmpl.ExecuteTemplate(w, "option_err", o)
	if err != nil {
		panic(err)
	}
}

func (o *Option) WriteValidation(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "option", o)
	if err != nil {
		panic(err)
	}
}
