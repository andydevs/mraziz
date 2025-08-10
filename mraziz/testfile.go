/* Bootstrap projects with created template */
package mraziz

import "time"

/* Test file data */
type TestfileData struct {
	TimeStamp string
	Name      string
	Adjective string
}

/* Generate test file */
func GetTestFile() (*File[TestfileData], error) {
	var err error

	// Get embedded template
	tmpl, err := GetEmbeddedTemplate("test")
	if err != nil {
		return nil, err
	}

	// Create file type
	filetype := File[TestfileData]{
		Template: tmpl,
		Filepath: "test.txt",
		Content: TestfileData{
			TimeStamp: time.Now().String(),
			Name:      "Parker",
			Adjective: "late",
		},
	}

	// Return filetype
	return &filetype, nil
}
