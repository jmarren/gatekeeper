package src

import (
	"fmt"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type TemplateWriter struct {
	ValidatorSpec *ValidatorSpec
	Field         *Field
	Value         any
}

func NewTemplateWriter(vSpec *ValidatorSpec, field *Field) *TemplateWriter {

	var data any
	switch vSpec.Name {
	case "max":
		data = util.AnyToInt(vSpec.Value, "value for max must be an int")
	case "min":
		data = util.AnyToInt(vSpec.Value, "value for min must be an int")
	case "minLen":
		data = util.AnyToInt(vSpec.Value, "value for minLen must be an int")
	case "maxLen":
		data = util.AnyToInt(vSpec.Value, "value for minLen must be an int")
	case "email":
		// import net/mail
		field.Obj.imports.Add(MAIL)
	case "option":
		data = util.AnyToStrSlice(vSpec.Value)
		field.Obj.imports.Add(SLICES)
	default:
		panic(fmt.Errorf("no validator named %s", vSpec.Name))
	}

	return &TemplateWriter{
		ValidatorSpec: vSpec,
		Field:         field,
		Value:         data,
	}
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
