package src

import (
	"io"

	"github.com/jmarren/gatekeeper/src/util"
)

type FieldSpec struct {
	Name            string
	Kind            string
	FormName        string           `yaml:"formName"`
	ValidationSpecs []*ValidatorSpec `yaml:"validators"`
	FmtKindErr      string           `yaml:"kindErr"`
}

func (f *FieldSpec) addImports(s util.StringSet) {

	for _, v := range f.ValidationSpecs {
		v.addImports(s)
	}

	if f.Kind == "int" {
		s.Add(STRCONV)
	}
}

func (f *FieldSpec) Validators(w io.Writer) []Validator {
	ret := []Validator{}
	for _, v := range f.ValidationSpecs {
		ret = append(ret, NewTemplateWriter(v, f, w))
	}
	return ret
}
