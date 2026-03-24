package src

import "text/template"

type Generator struct {
	templates *template.Template
	config    *Config
}

func NewGenerator(path string) *Generator {
	return &Generator{
		templates: BuildTemplates(),
		config:    NewConfig(path),
	}
}

func (g *Generator) Generate() {
	// generate each object using the provided templates
	for _, obj := range g.config.Objects {
		obj.Generate(g.templates)
	}
}
