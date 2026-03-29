package src

import (
	"os"
	"strings"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type Object struct {
	*ObjectSpec
	builder *strings.Builder
	Fields  []*Field
	imports util.StringSet
	Imports []string
}

func NewObject(spec *ObjectSpec) *Object {

	builder := new(strings.Builder)

	o := &Object{
		ObjectSpec: spec,
		builder:    builder,
		imports:    util.NewStringSet(),
		Fields:     []*Field{},
	}

	// add default imports
	o.imports.Add(HTTP)
	o.imports.Add(GATEKEEPER_ERR)

	for _, fs := range spec.FieldSpecs {
		field := NewField(fs, o)
		o.Fields = append(o.Fields, field)
	}

	o.Imports = o.imports.ToSlice()

	return o
}

func (o *Object) writeFields() {
	for _, f := range o.Fields {
		f.WriteValidation()
	}
}

func (o *Object) writeErrors() {
	for _, f := range o.Fields {
		f.WriteOuter()
	}
}

func (o *Object) writeFile() {

	// open the file
	// file := o.outFile()

	// defer closing
	// defer file.Close()

	path := o.outPath()

	err := os.WriteFile(path, []byte(o.builder.String()), 0777)

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
