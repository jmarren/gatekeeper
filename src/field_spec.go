package src

import (
	"io"
)

type FieldSpec struct {
	Name            string
	Kind            string
	FormName        string           `yaml:"formName"`
	ValidationSpecs []*ValidatorSpec `yaml:"validators"`
	FmtKindErr      string           `yaml:"kindErr"`
}

func (f *FieldSpec) Validators(w io.Writer) []Validator {
	ret := []Validator{}
	for _, v := range f.ValidationSpecs {
		ret = append(ret, NewTemplateWriter(v, f, w))
	}
	return ret
}
