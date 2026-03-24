package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jmarren/gatekeeper/src"
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of gatekeeper:\n")
	fmt.Fprintf(os.Stderr, "\tgatekeeper [path to yaml file]\n")
	flag.PrintDefaults()
}

func main() {
	// get path from args
	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}

	path := os.Args[1]

	generator := src.NewGenerator(path)

	generator.Generate()

}
