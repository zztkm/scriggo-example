package main

import (
	"log"

	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/native"
)

var packages native.Packages

func main() {

	// src is the source code of the program to run.
	src := []byte(`
	package main

	import (
		"fmt"
		"github.com/fatih/color"
	)

	func main() {
		fmt.Println("Here you are go")
		color.Red("Roses are red")
		color.Blue("Violets are blue")
	}
	`)

	// Create a file system with the file of the program to run.
	fsys := scriggo.Files{"main.go": src}

	// Use the importer in the packages variable.
	opts := &scriggo.BuildOptions{Packages: packages}

	// Build the program.
	program, err := scriggo.Build(fsys, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Run the program.
	err = program.Run(nil)
	if err != nil {
		log.Fatal(err)
	}

}
