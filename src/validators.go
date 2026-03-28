package src

import (
	"github.com/jmarren/gatekeeper/src/util"
)

type Validator interface {
	WriteErr()
	WriteValidation()
	imports() util.StringSet
}

type ValidatorSpec struct {
	Name   string
	Value  any
	FmtErr string `yaml:"error"`
}

func (v *ValidatorSpec) addImports(s util.StringSet) {
	switch v.Name {
	case "option":
		s.Add(SLICES)
	case "email":
		s.Add(MAIL)

	}
}

// func (v *ValidatorSpec) Validator(w io.Writer, f *FieldSpec) Validator {
// 	return NewTemplateWriter(v, f, w)
// }

// func (v *ValidatorSpec) WriteErr(field *FieldSpec, w io.Writer) {
// 	switch v.Name {
// 	case "minLen":
// 		minLen := NewMinLen(field, v)
// 		minLen.WriteError(w)
// 	case "maxLen":
// 		maxLen := NewMaxLen(field, v)
// 		maxLen.WriteError(w)
// 	case "email":
// 		email := NewEmail(field, v)
// 		email.WriteError(w)
// 	case "option":
// 		option := NewOption(field, v)
// 		option.WriteError(w)
// 	}
// }
//
// func (v *ValidatorSpec) WriteValidation(field *FieldSpec, w io.Writer) {
//
// 	switch v.Name {
// 	case "minLen":
// 		minLen := NewMinLen(field, v)
// 		minLen.WriteValidation(w)
// 	case "maxLen":
// 		maxLen := NewMaxLen(field, v)
// 		maxLen.WriteValidation(w)
// 	case "email":
// 		email := NewEmail(field, v)
// 		email.WriteValidation(w)
// 	case "option":
// 		option := NewOption(field, v)
// 		option.WriteValidation(w)
// 	}
// }
//
//	type Validators struct {
//		MinLen  string            `yaml:"minLen"`
//		MaxLen  string            `yaml:"maxLen"`
//		Min     string            `yaml:"min"`
//		Max     string            `yaml:"max"`
//		Options []string          `yaml:"options"`
//		Email   bool              `yaml:"email"`
//		Custom  []CustomValidator `yaml:"custom"`
//	}

// type CustomValidator struct {
// 	Package  string `yaml:"package"`
// 	Function string `yaml:"function"`
// }
