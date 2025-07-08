package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a hello message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Cobra CLI!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
