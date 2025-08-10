/* Bootstrap projects with created template */
package mraziz

/* Return base project configuration */
func GetBasePizza(project ProjectData) (*Directory, error) {
	var err error

	// Get test file template
	test, err := GetTestFile()
	if err != nil {
		return nil, err
	}

	// Get config file template
	confTmpl, err := GetEmbeddedTemplate("config.yaml")
	if err != nil {
		return nil, err
	}

	// Get config file
	conf := File[ProjectData]{
		Template: confTmpl,
		Filepath: "config.yaml",
		Content:  project,
	}

	// Get pizza directory
	pizza := Directory{
		Path:  ".pizza",
		Items: []Generatable{test, &conf},
	}

	// Return directory
	return &pizza, nil
}
