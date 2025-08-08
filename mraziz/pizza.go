package mraziz

import (
	"os"
)

func CreatePizzaDirectory() error {
	err := os.Mkdir(".pizza", 0755)
	return err
}
