package src

type Field struct {
	Name       string
	Kind       string
	FormName   string `yaml:"formName"`
	Validators `yaml:"validators,inline"`
}
