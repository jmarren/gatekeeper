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
	FmtKindErr string
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
	switch f.Kind {
	case "int":

	}
}

func (f *Field) WriteKindErrVar(w io.Writer) {
	templates.Tmpl.ExecuteTemplate(w, "kind_err_var", f)
}

func (f *Field) WriteValidation(w io.Writer) {
	f.WriteAssignment(w)
	for _, v := range f.Validators {
		v.WriteValidation(f, w)
	}
}

func (f *Field) WriteErrorVars(w io.Writer) {
	f.WriteKindErrVar(w)
	for _, v := range f.Validators {
		v.WriteErrVar(f, w)
	}
}

func (f *Field) WriteKindErrInit(w io.Writer) {

}

func (f *Field) WriteErrorInits(w io.Writer) {
	for _, v := range f.Validators {
		v.WriteErrInits(f, w)
	}
}

func (f *Field) init() {

	// set FormName to Name if not provided
	if f.FormName == "" {
		f.FormName = f.Name
	}
}
