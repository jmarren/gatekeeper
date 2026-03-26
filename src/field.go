package src

import (
	"fmt"
	"io"

	"github.com/jmarren/gatekeeper/src/util"
)

type Field struct {
	Name       string
	Kind       string
	FormName   string           `yaml:"formName"`
	Validators []*ValidatorSpec `yaml:"validators"`
	// Validators `yaml:"validators,inline"`
}

func (f *Field) addImports(s util.StringSet) {

	for _, v := range f.Validators {
		v.addImports(s)
	}

	if f.Kind == "int" {
		s.Add(STRCONV)
	}
}

// type Validator interface {
// 	Write(w io.Writer)
// }

func (f *Field) Write(w io.Writer) {
	for _, v := range f.Validators {
		v.Write(f, w)
	}
}

// func (f *Field) Validators() []Validator {
// 	for _, v := range f.Validators {
// 		swi
// 	}
// }

func (f *Field) init() {

	// set FormName to Name if not provided
	if f.FormName == "" {
		f.FormName = f.Name
	}

	for _, v := range f.Validators {
		if v.Name == "minLen" {
			fmt.Printf("minLen value = %v\n", v.Value)
		}
	}

	fmt.Printf("f.FormName = %s\n", f.FormName)
}
