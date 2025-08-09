/* Bootstrap projects with created template */
package mraziz

// Imports
import (
	"os"
	"text/template"
)

type File[D any] struct {
	Template *template.Template
	Filepath string
	Content  D
}

func (t *File[D]) Generate() error {
	// Open file for writing
	outf, err := os.Create(t.Filepath)
	if err != nil {
		return err
	}
	defer outf.Close()

	// Write data to file
	err = t.Template.Execute(outf, t.Content)
	return err
}
