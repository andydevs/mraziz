package mraziz

import (
	"time"
)

type TestfileData struct {
	TimeStamp string
	Name      string
	Adjective string
}

func getPizzaDirectory() *DirectoryType {
	template := DirectoryType{
		Path: ".pizza",
	}
	return &template
}

func getTestFile() (*FileType[TestfileData], error) {
	tmpl, err := GetEmbeddedTemplate("test")
	if err != nil {
		return nil, err
	}
	time := time.Now()
	filetype := FileType[TestfileData]{
		Template: tmpl,
		Filepath: "test.txt",
		Content: TestfileData{
			TimeStamp: time.String(),
			Name:      "Parker",
			Adjective: "late",
		},
	}
	return &filetype, nil
}

func GetBaseProject() (*DirectoryWithItemsType, error) {
	baseDir := getPizzaDirectory()
	template, err := getTestFile()
	if err != nil {
		return nil, err
	}
	ditems := DirectoryWithItemsType{
		Directory: *baseDir,
		Items: []Generatable{
			template,
		},
	}
	return &ditems, nil
}
