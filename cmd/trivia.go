/*
Copyright © 2025 NAME HERE dominicmaynard00@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/dom-m17/talk-to-me/utils"
	"github.com/spf13/cobra"
)

var (
	name  string
	times int
)

var greetCmd = &cobra.Command{
	Use:   "trivia",
	Short: "Prints a trivia question",
	Run:   getTrivia,
}

func init() {
	// Add the command to the root
	rootCmd.AddCommand(greetCmd)

	// Local flags (only for this command)
	// TODO: use a flag here that determines if the question will be multiple choice or true/false
	// greetCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
}

func getTrivia(cmd *cobra.Command, args []string) {
	cfg, err := ReadConfig()
	if err != nil {
		fmt.Println("reading config", err)
		return
	}

	var fact Trivia
	fact, err = utils.GetDataFromAPI(
		cfg.OpenTriviaBaseURL.String(),
		"",
		fact,
	)

	fmt.Println(fact.Results[0].Question)

}
