package src

import (
	"os"
	"strings"
	"text/template"

	"github.com/jmarren/gatekeeper/src/templates"
	"github.com/jmarren/gatekeeper/src/util"
)

func BuildTemplates() *template.Template {

	tmpl := template.New("base").Funcs(template.FuncMap{
		"joinStrs": util.JoinStrings,
		"join":     strings.Join,
	})

	templatesPath := templates.DirPath()

	entries, err := os.ReadDir(templatesPath)

	if err != nil {
		panic(err)
	}

	// add all files with .tmpl extension to templates
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".tmpl") {
			tmpl = template.Must(tmpl.ParseFiles(templatesPath + name))
		}
	}

	return tmpl

}
