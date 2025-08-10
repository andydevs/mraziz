/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/andydevs/mraziz/mraziz"
	"github.com/spf13/cobra"
)

func executeInit(cmd *cobra.Command, args []string) {
	var err error

	// Make sure we have correct args
	if len(args) == 0 {
		err = fmt.Errorf("please provide a template path or url")
		panic(err)
	}

	// Print args
	pathOrUrlArg := args[0]
	fmt.Printf("Path/URL being used: %s\n", pathOrUrlArg)

	// Get directory name for project
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirname := path.Base(dir)

	// Create project config
	project := mraziz.ProjectData{
		ProjectTitle:       dirname,
		ProjectDescription: dirname,
		ProjectTemplate:    pathOrUrlArg,
	}

	// Create pizza directory
	pizza, err := mraziz.GetBasePizza(project)
	if err != nil {
		panic(err)
	}

	// Generate
	err = pizza.Generate()
	if err != nil {
		panic(err)
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your new project",
	Long:  `Creates a .pizza directory and run some initialization`,
	Run:   executeInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
