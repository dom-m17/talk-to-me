/*
Copyright © 2025 NAME HERE dominicmaynard00@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	name  string
	times int
)

var greetCmd = &cobra.Command{
	Use:   "trivia",
	Short: "Prints a trivia question",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// Add the command to the root
	rootCmd.AddCommand(greetCmd)

	// Local flags (only for this command)
	// TODO: use a flag here that determines if the question will be multiple choice or true/false
	// greetCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
}
