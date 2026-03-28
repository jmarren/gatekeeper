package src

import (
	"os"
	"strings"
)

type ObjectSpec struct {
	Name       string       `yaml:"name"`
	Package    string       `yaml:"package"`
	FieldSpecs []*FieldSpec `yaml:"fields"`
	Path       string       `yaml:"path"`
	Imports    []string
}

// generate the outPath for this object
func (o *ObjectSpec) outPath() string {
	// allow for {path}, /{path}, and ./{path} formats
	path, _ := strings.CutSuffix(o.Path, "/")
	path, _ = strings.CutPrefix(path, "./")
	path, _ = strings.CutPrefix(path, "/")
	// concatenate .gatekeeper.go to file name
	return path + "/" + o.Name + ".gatekeeper.go"
}

// open the outfile for this object and return it
func (o *ObjectSpec) outFile() *os.File {
	file, err := os.OpenFile(o.outPath(), os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		panic(err)
	}
	return file
}

// // range over fields and add any required imports to the imports set
// func (o *ObjectSpec) setImports() {
// 	imports := util.NewStringSet()
// 	imports.Add(HTTP)
// 	imports.Add(GATEKEEPER_ERR)
// 	for _, field := range o.FieldSpecs {
// 		field.addImports(imports)
// 	}
// 	o.Imports = imports.ToSlice()
// }
//
// // Generate the .gatekeeper.go file for this object
// func (o *ObjectSpec) Generate() {
//
// 	o.setImports()
//
// 	file := o.outFile()
// 	defer file.Close()
//
// 	err := templates.Tmpl.ExecuteTemplate(file, "header", o)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	err = templates.Tmpl.ExecuteTemplate(file, "typedef", o)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	o.WriteErrors(file)
//
// 	// o.WriteErrorInits(file)
//
// 	// _, err = file.WriteString("}\n")
// 	//
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	//
// 	err = templates.Tmpl.ExecuteTemplate(file, "constructor", o)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	o.WriteFields(file)
//
// 	_, err = file.WriteString("\n\treturn x, errGroup \n}")
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	err = templates.Tmpl.ExecuteTemplate(file, "handler", o)
//
// 	if err != nil {
// 		panic(err)
// 	}
// }
