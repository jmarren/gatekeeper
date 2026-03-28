package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type Field struct {
	w io.Writer
	*FieldSpec
	Validators []Validator
	imports    util.StringSet
}

func NewField(s *FieldSpec, w io.Writer) *Field {

	// set FormName to Name if not provided
	if s.FormName == "" {
		s.FormName = s.Name
	}

	// set default KindErrs
	if s.FmtKindErr == "" {
		if s.Kind == "int" {
			s.FmtKindErr = s.FormName + " must be an int"
		}
		if s.Kind == "string" {
			s.FmtKindErr = s.FormName + " must be a string"
		}
	}

	// create validators from fieldSpec
	validators := s.Validators(w)

	// create imports set
	imports := util.NewStringSet()

	// merge in all validators imports
	for _, v := range validators {
		imports.Merge(v.imports())
	}

	return &Field{
		w:          w,
		FieldSpec:  s,
		Validators: validators,
		imports:    imports,
	}
}

func (f *Field) WriteValidation() {
	f.WriteAssignment()

	for _, v := range f.Validators {
		v.WriteValidation()
	}
}

func (f *Field) WriteAssignment() {
	var err error
	switch f.Kind {
	case "int":
		err = templates.Tmpl.ExecuteTemplate(f.w, "int", f)
	case "string":
		err = templates.Tmpl.ExecuteTemplate(f.w, "string", f)
	default:
		panic("kind must be string or int")
	}

	util.PanicIf(err)
}

func (f *Field) WriteErrors() {
	err := templates.Tmpl.ExecuteTemplate(f.w, "kind_err", f)
	util.PanicIf(err)
	for _, v := range f.Validators {
		v.WriteErr()
	}
}
