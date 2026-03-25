package src

import (
	"os"
	"strings"
	"text/template"

	"github.com/jmarren/gatekeeper/src/util"
)

type Object struct {
	Name    string   `yaml:"name"`
	Package string   `yaml:"package"`
	Fields  []*Field `yaml:"fields"`
	Path    string   `yaml:"path"`
	Imports []string
}

func (o *Object) outPath() string {
	path, _ := strings.CutSuffix(o.Path, "/")
	path, _ = strings.CutPrefix(path, "./")
	path, _ = strings.CutPrefix(path, "/")
	return path + "/" + o.Name + ".gatekeeper.go"
}

func (o *Object) outFile() *os.File {
	file, err := os.OpenFile(o.outPath(), os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		panic(err)
	}
	return file
}

func (o *Object) init() {
	for _, field := range o.Fields {
		field.init()
	}
}

func (o *Object) setImports() {
	imports := util.NewStringSet()
	imports.Add(HTTP)
	imports.Add(FMT)
	for _, field := range o.Fields {
		field.addImports(imports)
	}
	o.Imports = imports.ToSlice()
}

func (o *Object) Generate(tmpl *template.Template) {

	o.init()
	o.setImports()

	file := o.outFile()
	defer file.Close()

	err := tmpl.ExecuteTemplate(file, "base", o)

	if err != nil {
		panic(err)
	}
}
