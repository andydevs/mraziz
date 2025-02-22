package main

import (
	"embed"
	"os"
	"text/template"
)

type Person struct {
	Name      string
	Adjective string
}

//go:embed res/*
var efs embed.FS

func main() {
	// Read embedded file
	tmplfileBytes, err := efs.ReadFile("res/test.tmpl")
	if err != nil {
		panic(err)
	}
	tmplFileText := string(tmplfileBytes)

	// Create template
	tmpl, err := template.New("test.tmpl").Parse(tmplFileText)
	if err != nil {
		panic(err)
	}

	// Print examples
	people := []Person{
		{"Parker", "late"},
		{"Megan", "high"},
		{"Gordon", "squishy"},
		{"Harry", "hairy"},
		{"Dan", "the man"},
	}
	for _, person := range people {
		err = tmpl.Execute(os.Stdout, person)
		if err != nil {
			panic(err)
		}
	}

}
