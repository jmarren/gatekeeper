package src

import (
	"strings"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type Object struct {
	*ObjectSpec
	builder *strings.Builder
	fields  []*Field
	imports util.StringSet
}

func NewObject(spec *ObjectSpec) *Object {

	fields := []*Field{}

	builder := new(strings.Builder)

	for _, fs := range spec.FieldSpecs {
		field := NewField(fs, builder)
		fields = append(fields, field)
	}

	imports := util.NewStringSet()
	imports.Add(HTTP)
	imports.Add(GATEKEEPER_ERR)

	// merge imports from all fields
	for _, field := range fields {
		imports.Merge(field.imports)
	}

	return &Object{
		ObjectSpec: spec,
		fields:     fields,
		builder:    builder,
		imports:    imports,
	}

}

func (o *Object) writeFields() {
	for _, f := range o.fields {
		f.WriteValidation()
	}
}

func (o *Object) writeErrors() {
	for _, f := range o.fields {
		f.WriteErrors()
	}
}

func (o *Object) writeFile() {

	// open the file
	file := o.outFile()

	// defer closing
	defer file.Close()

	// write the string builders string to the file
	_, err := file.WriteString(o.builder.String())

	util.PanicIf(err)

}

func (o *Object) Write() {
	var err error

	// TODO: write header
	err = templates.Tmpl.ExecuteTemplate(o.builder, "header", o)
	util.PanicIf(err)

	// write errors
	o.writeErrors()

	// write type definition
	err = templates.Tmpl.ExecuteTemplate(o.builder, "typedef", o)
	util.PanicIf(err)

	// write open constructor
	err = templates.Tmpl.ExecuteTemplate(o.builder, "constructor", o)
	util.PanicIf(err)

	// write validation for each field
	o.writeFields()

	// write close constructor
	_, err = o.builder.WriteString("\n\treturn x, errGroup \n}")
	util.PanicIf(err)

	// write handler
	err = templates.Tmpl.ExecuteTemplate(o.builder, "handler", o)
	util.PanicIf(err)

	// write the builder to file
	o.writeFile()

}
