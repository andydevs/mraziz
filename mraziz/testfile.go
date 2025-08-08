package mraziz

import (
	"os"
)

type TestFile struct {
	Name      string
	Adjective string
}

func (t *TestFile) CreateFile() {
	// Get template
	tmpl, err := GetEmbeddedTemplate("test")
	if err != nil {
		panic(err)
	}

	// Open file for writing
	outf, err := os.Create(".pizza/test.txt")
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	// Write data to file
	err = tmpl.Execute(outf, t)
	if err != nil {
		panic(err)
	}
}
