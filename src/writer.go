package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type Importer interface {
	imports() util.StringSet
}

type TemplateWriter struct {
	w     io.Writer
	vSpec *ValidatorSpec
	data  Importer
}

func (t *TemplateWriter) imports() util.StringSet {
	return t.data.imports()
}

func NewTemplateWriter(vSpec *ValidatorSpec, field *FieldSpec, w io.Writer) *TemplateWriter {

	t := &TemplateWriter{
		w:     w,
		vSpec: vSpec,
	}

	switch vSpec.Name {
	case "minLen":
		t.data = NewMinLen(field, vSpec)
	case "maxLen":
		t.data = NewMaxLen(field, vSpec)
	}

	return t
}

func (t *TemplateWriter) errTemplateName() string {
	return t.vSpec.Name + "_err"
}

func (t *TemplateWriter) WriteErr() {
	err := templates.Tmpl.ExecuteTemplate(t.w, t.errTemplateName(), t.data)
	util.PanicIf(err)
}

func (t *TemplateWriter) WriteValidation() {
	err := templates.Tmpl.ExecuteTemplate(t.w, t.vSpec.Name, t.data)
	util.PanicIf(err)
}
