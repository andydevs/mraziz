package mraziz

import (
	"embed"
	gotmpl "text/template"
)

//go:embed res/*
var efs embed.FS

func GetEmbeddedTemplate(name string) (*gotmpl.Template, error) {
	filename := "res/" + name + ".tmpl"

	// Read embedded file
	tmplfileBytes, err := efs.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	tmplFileText := string(tmplfileBytes)

	// Create template
	tmpl, err := gotmpl.New(filename).Parse(tmplFileText)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
