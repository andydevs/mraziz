/* Bootstrap projects with created template */
package mraziz

/* Return base project configuration */
func GetBasePizza() (*Directory, error) {
	var err error

	// Get test file template
	test, err := GetTestFile()
	if err != nil {
		return nil, err
	}

	// Get pizza directory
	pizza := Directory{
		Path:  ".pizza",
		Items: []Generatable{test},
	}

	// Return directory
	return &pizza, nil
}
