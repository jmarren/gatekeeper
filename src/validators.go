package src

type CustomValidator struct {
	Package  string `yaml:"package"`
	Function string `yaml:"function"`
}

type Validators struct {
	MinLen  string            `yaml:"minLen"`
	MaxLen  string            `yaml:"maxLen"`
	Min     string            `yaml:"min"`
	Max     string            `yaml:"max"`
	Options []string          `yaml:"options"`
	Email   bool              `yaml:"email"`
	Custom  []CustomValidator `yaml:"custom"`
}
