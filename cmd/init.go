/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/andydevs/mraziz/pizza"
	"github.com/spf13/cobra"
)

// For test file
var name string
var adjective string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your new project",
	Long:  `Creates a .pizza directory and run some initialization`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create .pizza directory

		// Create test file
		pizza.CreateTestFile(name, adjective)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().StringVar(&name, "test-name", "Parker", "name to write to test file")
	initCmd.Flags().StringVar(&adjective, "test-adj", "late", "adjective to write to test file")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
