package src

import (
	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type TemplateWriter struct {
	ValidatorSpec *ValidatorSpec
	Field         *Field
	Value         any
}

type validator struct {
	converter func(t *TemplateWriter)
	imports   []string
}

func ValToInt(t *TemplateWriter) {
	val, ok := t.ValidatorSpec.Value.(int)

	if !ok {
		panic("value for " + t.ValidatorSpec.Name + " must be an int")
	}
	t.Value = val
}

func ValToStringArr(t *TemplateWriter) {

	vals := []string{}

	iVals, ok := t.ValidatorSpec.Value.([]any)
	if !ok {
		panic(t.ValidatorSpec.Name + " value must be a list")
	}

	for _, iVal := range iVals {
		val, ok := iVal.(string)
		if !ok {
			panic(t.ValidatorSpec.Name + " value must be a list of strings")
		}
		vals = append(vals, val)

	}

	t.Value = vals
}

func EmptyConverter(t *TemplateWriter) {}

func NewValidator() *validator {
	return &validator{
		imports:   []string{},
		converter: EmptyConverter,
	}
}

func (v *validator) Import(i ...string) *validator {
	v.imports = append(v.imports, i...)
	return v
}

func (v *validator) Use(fn func(t *TemplateWriter)) *validator {
	curr := v.converter
	v.converter = func(t *TemplateWriter) {
		curr(t)
		fn(t)
	}
	return v
}

var simpleIntValidator *validator = NewValidator().Use(ValToInt)

var validators map[string]*validator = map[string]*validator{
	"max":    simpleIntValidator,
	"min":    simpleIntValidator,
	"maxLen": simpleIntValidator,
	"minLen": simpleIntValidator,
	"email":  NewValidator().Import(MAIL),
	"option": NewValidator().Use(ValToStringArr).Import(STRCONV),
}

func NewTemplateWriter(vSpec *ValidatorSpec, field *Field) *TemplateWriter {

	validator, ok := validators[vSpec.Name]

	if !ok {
		panic("no validator named " + vSpec.Name)
	}

	t := &TemplateWriter{
		ValidatorSpec: vSpec,
		Field:         field,
	}

	validator.converter(t)

	field.Obj.imports.Add(validator.imports...)

	return t
}

func (t *TemplateWriter) errTemplateName() string {
	return t.ValidatorSpec.Name + "_err"
}

func (t *TemplateWriter) WriteOuter() {
	err := templates.Tmpl.ExecuteTemplate(t.Field.Obj.builder, t.errTemplateName(), t)
	util.PanicIf(err)
}

func (t *TemplateWriter) WriteValidation() {
	err := templates.Tmpl.ExecuteTemplate(t.Field.Obj.builder, t.ValidatorSpec.Name, t)
	util.PanicIf(err)
}
