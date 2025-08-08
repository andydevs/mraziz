/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/andydevs/mraziz/mraziz"
	"github.com/spf13/cobra"
)

func executeInit(cmd *cobra.Command, args []string) {
	var err error

	// Create pizza directory
	project, err := mraziz.GetBaseProject()
	if err != nil {
		panic(err)
	}

	// Generate
	err = project.Generate()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
