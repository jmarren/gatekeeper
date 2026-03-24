package src

import (
	"os"
	"slices"
	"strings"
	"text/template"
)

type Object struct {
	Name    string   `yaml:"name"`
	Package string   `yaml:"package"`
	Fields  []*Field `yaml:"fields"`
	Path    string   `yaml:"path"`
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

func (o *Object) imports() []string {
	imports := []string{"\"net/http\"", "\"fmt\""}
	for _, field := range o.Fields {
		if slices.Contains() {
		}
	}
	return imports
}

func (o *Object) Generate(tmpl *template.Template) {

	o.init()

	file := o.outFile()
	defer file.Close()

	err := tmpl.ExecuteTemplate(file, "base", o)
	if err != nil {
		panic(err)
	}
}
