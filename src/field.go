package src

import (
	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type Field struct {
	obj *Object
	*FieldSpec
	Validators []Validator
}

func NewField(spec *FieldSpec, obj *Object) *Field {

	// set FormName to Name if not provided
	if spec.FormName == "" {
		spec.FormName = spec.Name
	}

	// set default KindErrs
	if spec.FmtKindErr == "" {
		if spec.Kind == "int" {
			spec.FmtKindErr = spec.FormName + " must be an int"
		}
		if spec.Kind == "string" {
			spec.FmtKindErr = spec.FormName + " must be a string"
		}
	}

	// create validators from fieldSpec
	validators := spec.Validators(obj.builder)

	f := &Field{
		obj:        obj,
		FieldSpec:  spec,
		Validators: []Validator{},
	}

	for _, v := range spec.ValidationSpecs {
		f.Validators = append(f.Validators, NewTemplateWriter(v, f))
	}

	return &Field{
		obj:        obj,
		FieldSpec:  spec,
		Validators: validators,
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
		err = templates.Tmpl.ExecuteTemplate(f.obj.builder, "int", f)
	case "string":
		err = templates.Tmpl.ExecuteTemplate(f.obj.builder, "string", f)
	default:
		panic("kind must be string or int")
	}

	util.PanicIf(err)
}

func (f *Field) WriteErrors() {
	err := templates.Tmpl.ExecuteTemplate(f.obj.builder, "kind_err", f)
	util.PanicIf(err)
	for _, v := range f.Validators {
		v.WriteErr()
	}
}
