package src

import (
	"strings"
)

type ObjectSpec struct {
	Name               string          `yaml:"name"`
	Package            string          `yaml:"package"`
	FieldSpecs         []*FieldSpec    `yaml:"fields"`
	Path               string          `yaml:"path"`
	EmitInterfacePaths []InterfacePath `yaml:"emit_interface_to"`
}

// generate the outPath for this object
func (o *ObjectSpec) outPath() string {
	// allow for {path}, /{path}, and ./{path} formats
	path, _ := strings.CutSuffix(o.Path, "/")
	path, _ = strings.CutPrefix(path, "./")
	path, _ = strings.CutPrefix(path, "/")
	// concatenate .gatekeeper.go to file name
	return path + "/" + o.Name + ".gatekeeper.go"
}
