package mraziz

type TestfileData struct {
	Name      string
	Adjective string
}

func GetTestFile(content TestfileData) (*FileType[TestfileData], error) {
	tmpl, err := GetEmbeddedTemplate("test")
	if err != nil {
		return nil, err
	}
	filetype := FileType[TestfileData]{
		Template: tmpl,
		Filepath: ".pizza/test.txt",
		Content:  content,
	}
	return &filetype, nil
}
