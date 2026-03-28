package src

import (
	"fmt"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type TemplateWriter struct {
	vSpec *ValidatorSpec
	field *Field
	data  any
}

func NewTemplateWriter(vSpec *ValidatorSpec, field *Field) *TemplateWriter {

	fmt.Println("NewTemplateWriter")

	var data any
	switch vSpec.Name {
	case "minLen":
		data = NewMinLen(field, vSpec)
	case "maxLen":
		data = NewMaxLen(field, vSpec)
	default:
		panic(fmt.Errorf("no validator named %s", vSpec.Name))
	}

	return &TemplateWriter{
		vSpec,
		field,
		data,
	}
}

func (t *TemplateWriter) errTemplateName() string {
	return t.vSpec.Name + "_err"
}

func (t *TemplateWriter) WriteErr() {
	err := templates.Tmpl.ExecuteTemplate(t.field.obj.builder, t.errTemplateName(), t.data)
	util.PanicIf(err)
}

func (t *TemplateWriter) WriteValidation() {
	err := templates.Tmpl.ExecuteTemplate(t.field.obj.builder, t.vSpec.Name, t.data)
	util.PanicIf(err)
}
