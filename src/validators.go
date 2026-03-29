package src

type Validator interface {
	WriteOuter()
	WriteValidation()
}

type ValidatorSpec struct {
	Name   string
	Value  any
	FmtErr string `yaml:"error"`
}
