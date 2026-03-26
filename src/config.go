package src

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Objects []*Object `yaml:"objects"`
}

func NewConfig(path string) *Config {
	yamlBytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	cfg := Config{}

	err = yaml.Unmarshal(yamlBytes, &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}

func (c *Config) Generate() {
	for _, obj := range c.Objects {
		obj.Generate()
	}
}
