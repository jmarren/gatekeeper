package yaml

import (
	"fmt"
	"os"
)

func TryNode() {
	yamlBytes, err := os.ReadFile("/home/john-marren/code/proj/gatekeeper/src/yaml/example.yaml")

	if err != nil {
		panic(err)
	}

	fmt.Printf("example.yaml = %s\n", string(yamlBytes))

	// cfg := Config{}
	//
	// err = yaml.Unmarshal(yamlBytes, &cfg)
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Printf("cfg = %v\n", cfg)
	//
	// return &cfg
}
