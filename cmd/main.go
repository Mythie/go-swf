package main

import (
	"os"
	"strings"

	"github.com/mythie/go-swf/pkg/swfparser"
)

func main() {
	println("Hello, world")

	// gets the current working directory
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	f, err := os.Open(strings.Join([]string{cwd, "fixtures", "air.swf"}, "/"))

	if err != nil {
		panic(err)
	}

	_, err = swfparser.FromReader(f)

	if err != nil {
		panic(err)
	}
}
