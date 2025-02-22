package pizza

import (
	"os"
)

type Person struct {
	Name      string
	Adjective string
}

func CreateTestFile(name string, adjective string) {
	// Build data structure
	person := Person{name, adjective}

	// Get template
	tmpl, err := GetEmbeddedTemplate("test")
	if err != nil {
		panic(err)
	}

	// Open file for writing
	outf, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	err = tmpl.Execute(outf, person)
	if err != nil {
		panic(err)
	}
}
