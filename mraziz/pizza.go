package mraziz

import (
	"os"
)

func CreatePizzaDirectory() {
	err := os.Mkdir(".pizza", 0755)
	if err != nil {
		panic(err)
	}
}
