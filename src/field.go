package src

import (
	"fmt"

	"github.com/jmarren/gatekeeper/src/util"
)

type Field struct {
	Name       string
	Kind       string
	FormName   string `yaml:"formName"`
	Validators `yaml:"validators,inline"`
}

func (f *Field) addImports(s *util.StringSet) {
	if f.Email {
		s.Add(MAIL)
	}
	if len(f.Options) > 0 {
		s.Add(SLICES)
	}

	if f.Kind == "int" {
		s.Add(STRCONV)
	}

}

func (f *Field) init() {

	fmt.Printf("f.FormName = %s\n", f.FormName)

	// set FormName to Name if not provided
	if f.FormName == "" {
		f.FormName = f.Name
	}

	fmt.Printf("f.FormName = %s\n", f.FormName)
}
