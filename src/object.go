package src

import (
	"os"
	"text/template"
)

type Object struct {
	Name    string  `yaml:"name"`
	Package string  `yaml:"package"`
	Fields  []Field `yaml:"fields"`
	Path    string  `yaml:"path"`
}

func (o *Object) outPath() string {
	return o.Path + "/" + o.Name + ".gatekeeper.go"
}

func (o *Object) Generate(tmpl *template.Template) {

	file, err := os.OpenFile(o.outPath(), os.O_WRONLY|os.O_CREATE, 0777)

	err = tmpl.ExecuteTemplate(file, "base", o)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
