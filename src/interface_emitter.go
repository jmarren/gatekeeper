package src

import (
	"os"
	"strings"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

type InterfacePath struct {
	Path    string
	Package string
}

type InterfaceWriter struct {
	InterfacePath
	Object
}

func NewInterfaceWriter(i InterfacePath, o Object) InterfaceWriter {
	return InterfaceWriter{
		i,
		o,
	}
}

func (i InterfaceWriter) Write() {

	builder := new(strings.Builder)
	err := templates.Tmpl.ExecuteTemplate(builder, "interface", i)

	util.PanicIf(err)
	err = os.WriteFile(i.Path+"/"+i.Name+".gatekeeper.interface.go", []byte(builder.String()), 0777)

	util.PanicIf(err)
}
