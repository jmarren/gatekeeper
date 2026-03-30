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

func (o *Object) Import(imports ...string) {
	o.imports.Add(imports...)
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
	o.Import(HTTP, GATEKEEPER_ERR)

	// create fields from field specs
	for _, fs := range spec.FieldSpecs {
		o.Fields = append(o.Fields, NewField(fs, o))
	}

	// set imports slice
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

	path := o.outPath()

	err := os.WriteFile(path, []byte(o.builder.String()), 0777)

	util.PanicIf(err)

}

func (o *Object) execTemplate(name string) {
	err := templates.Tmpl.ExecuteTemplate(o.builder, name, o)
	util.PanicIf(err)
}

func (o *Object) Write() {
	var err error

	// write header
	o.execTemplate("header")

	// write errors
	o.writeErrors()

	// write type definition
	o.execTemplate("typedef")

	// write open constructor
	o.execTemplate("constructor")

	// write validation for each field
	o.writeFields()

	// write close constructor
	_, err = o.builder.WriteString("\n\treturn x, errGroup \n}")
	util.PanicIf(err)

	// write handler
	o.execTemplate("handler")

	// write the builder to file
	o.writeFile()

}
