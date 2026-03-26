package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type Field struct {
	Name       string
	Kind       string
	FormName   string           `yaml:"formName"`
	Validators []*ValidatorSpec `yaml:"validators"`
	FmtKindErr string           `yaml:"kindErr"`
}

func (f *Field) addImports(s util.StringSet) {

	for _, v := range f.Validators {
		v.addImports(s)
	}

	if f.Kind == "int" {
		s.Add(STRCONV)
	}
}

func (f *Field) WriteAssignment(w io.Writer) {
	var err error
	switch f.Kind {
	case "int":
		err = templates.Tmpl.ExecuteTemplate(w, "int", f)
	case "string":
		err = templates.Tmpl.ExecuteTemplate(w, "string", f)
	}

	if err != nil {
		panic(err)
	}

}

func (f *Field) WriteValidation(w io.Writer) {
	f.WriteAssignment(w)
	for _, v := range f.Validators {
		v.WriteValidation(f, w)
	}
}

func (f *Field) WriteErrors(w io.Writer) {
	f.WriteKindErr(w)
	for _, v := range f.Validators {
		v.WriteErr(f, w)
	}
}

func (f *Field) WriteKindErr(w io.Writer) {
	err := templates.Tmpl.ExecuteTemplate(w, "kind_err", f)
	if err != nil {
		panic(err)
	}
}

func (f *Field) init() {

	// set FormName to Name if not provided
	if f.FormName == "" {
		f.FormName = f.Name
	}

	if f.FmtKindErr == "" {
		f.FmtKindErr = f.FormName + " must be "
		if f.Kind == "int" {
			f.FmtKindErr += "an int"
		}

		if f.Kind == "string" {
			f.FmtKindErr += "a string"
		}
	}

}
