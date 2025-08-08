package mraziz

import "os"

type DirectoryType struct {
	Path string
}

func (t *DirectoryType) Generate() error {
	err := os.Mkdir(t.Path, 0755)
	return err
}

type DirectoryWithItemsType struct {
	Directory DirectoryType
	Items     []Generatable
}

func (t *DirectoryWithItemsType) Generate() error {
	var err error

	err = t.Directory.Generate()
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Chdir(t.Directory.Path)
	if err != nil {
		return err
	}
	defer os.Chdir(dir)

	for _, item := range t.Items {
		err = item.Generate()
		if err != nil {
			return err
		}
	}

	return nil
}
