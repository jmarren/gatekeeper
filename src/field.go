package src

import (
	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type Field struct {
	*FieldSpec
	Obj        *Object
	Validators []Validator
}

func NewField(spec *FieldSpec, obj *Object) *Field {

	// set FormName to Name if not provided
	if spec.FormName == "" {
		spec.FormName = spec.Name
	}

	if spec.Kind == "int" {
		obj.imports.Add(STRCONV)
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

	f := &Field{
		Obj:        obj,
		FieldSpec:  spec,
		Validators: []Validator{},
	}

	for _, v := range spec.ValidationSpecs {
		f.Validators = append(f.Validators, NewTemplateWriter(v, f))
	}

	return f
}

func (f *Field) WriteValidation() {
	var err error
	_, err = f.Obj.builder.WriteString("\n	// " + f.Name + " validation")
	util.PanicIf(err)

	f.WriteAssignment()

	for _, v := range f.Validators {
		v.WriteValidation()
	}
}

func (f *Field) WriteAssignment() {
	var err error
	switch f.Kind {
	case "int":
		err = templates.Tmpl.ExecuteTemplate(f.Obj.builder, "int", f)
	case "string":
		err = templates.Tmpl.ExecuteTemplate(f.Obj.builder, "string", f)
	default:
		panic("kind must be string or int")
	}

	util.PanicIf(err)
}

func (f *Field) WriteOuter() {
	var err error

	_, err = f.Obj.builder.WriteString("// " + f.Name + " errors")
	util.PanicIf(err)

	err = templates.Tmpl.ExecuteTemplate(f.Obj.builder, "kind_err", f)
	util.PanicIf(err)
	for _, v := range f.Validators {
		v.WriteOuter()
	}
}
