package src

type FieldSpec struct {
	Name            string
	Kind            string
	FormName        string           `yaml:"formName"`
	ValidationSpecs []*ValidatorSpec `yaml:"validators"`
	FmtKindErr      string           `yaml:"kindErr"`
}
