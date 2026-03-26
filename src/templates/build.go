package templates

import (
	"os"
	"strings"
	"text/template"

	"github.com/jmarren/gatekeeper/src/util"
)

var Tmpl *template.Template

func init() {
	Tmpl = template.New("base").Funcs(template.FuncMap{
		"joinStrs": util.JoinStrings,
		"join":     strings.Join,
	})

	templatesPath := dirPath()

	entries, err := os.ReadDir(templatesPath)

	if err != nil {
		panic(err)
	}

	// add all files with .tmpl extension to templates
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".tmpl") {
			Tmpl = template.Must(Tmpl.ParseFiles(templatesPath + name))
		}
	}

}
