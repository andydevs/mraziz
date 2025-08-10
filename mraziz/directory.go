/* Bootstrap projects with created template */
package mraziz

// Imports
import "os"

/* Generate a directory with included items */
type Directory struct {
	Path  string
	Items []Generatable
}

/* Generate function for directory */
func (t *Directory) Generate() error {
	var err error

	// Get current directory to change back to
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Create directory
	err = os.Mkdir(t.Path, 0755)
	if err != nil {
		return err
	}

	// Change into current directory (change back at end)
	err = os.Chdir(t.Path)
	if err != nil {
		return err
	}
	defer os.Chdir(dir)

	// Create items
	for _, item := range t.Items {
		err = item.Generate()
		if err != nil {
			return err
		}
	}

	// End
	return nil
}
