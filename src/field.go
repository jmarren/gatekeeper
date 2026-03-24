package src

import "fmt"

type Field struct {
	Name       string
	Kind       string
	FormName   string `yaml:"formName"`
	Validators `yaml:"validators,inline"`
}

func (f *Field) imports() []string {
	imports := []string{}
	if f.Email {
		imports = append(imports, "\"net/email\"")
	}
	if len(f.Options) > 0 {
		imports = append(imports, "\"slices\"")
	}

	if f.Kind == "int" {
		imports = append(imports, "\"strconv\"")
	}

	return imports
}

func (f *Field) init() {

	fmt.Printf("f.FormName = %s\n", f.FormName)

	// set FormName to Name if not provided
	if f.FormName == "" {
		f.FormName = f.Name
	}

	fmt.Printf("f.FormName = %s\n", f.FormName)
}
