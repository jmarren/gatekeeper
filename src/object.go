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

	builder := new(strings.Builder)

	o := &Object{
		ObjectSpec: spec,
		builder:    builder,
		imports:    util.NewStringSet(),
		fields:     []*Field{},
	}

	// add default imports
	o.imports.Add(HTTP)
	o.imports.Add(GATEKEEPER_ERR)

	for _, fs := range spec.FieldSpecs {
		field := NewField(fs, o)
		o.fields = append(o.fields, field)
	}

	return o
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
