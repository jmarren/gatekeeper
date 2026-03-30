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

func SimpleConverter(t *TemplateWriter) {
	val, ok := t.ValidatorSpec.Value.(string)
	if !ok {
		panic("value for " + t.ValidatorSpec.Name + " must be a string")
	}
	t.Value = val
}

func (v *validator) UseErrVar() *validator {
	v.Use(func(t *TemplateWriter) {
		t.Field.Obj.UseErrVar()
	})
	return v
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
	"email":  NewValidator().Import(MAIL).UseErrVar(),
	"option": NewValidator().Use(ValToStringArr).Import(STRCONV, SLICES).UseErrVar(),
	"regex":  NewValidator().Use(SimpleConverter).Import(REGEX).UseErrVar(),
}

func NewTemplateWriter(vSpec *ValidatorSpec, field *Field) *TemplateWriter {

	// get the validator from the lookup table
	validator, ok := validators[vSpec.Name]

	// panic if not found
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
