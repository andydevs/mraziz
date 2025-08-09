/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path"

	"github.com/andydevs/mraziz/mraziz"
	"github.com/spf13/cobra"
)

func executeInit(cmd *cobra.Command, args []string) {
	var err error

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
