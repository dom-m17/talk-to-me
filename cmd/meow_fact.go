/*
Copyright © 2025 NAME HERE dominicmaynard00@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type MeowFact struct {
	Data []string `json:"data"`
}

// meowFactCmd represents the meowFact command
var meowFactCmd = &cobra.Command{
	Use:     "meowFact",
	Aliases: []string{"meow-fact"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := ReadConfig()
		if err != nil {
			fmt.Println("reading config", err)
			return
		}

		client := &http.Client{}
		request, err := http.NewRequest(
			http.MethodGet,
			cfg.MeowFactBaseURL.String(),
			nil,
		)
		if err != nil {
			fmt.Println("create error:", err)
			return
		}

		resp, err := client.Do(request)
		if err != nil {
			fmt.Println("get error")
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		var fact MeowFact
		err = json.Unmarshal(body, &fact)
		if err != nil {
			fmt.Println("unmarshal error:", err)
		}

		fmt.Println(fact.Data[0])
	},
}

func init() {
	rootCmd.AddCommand(meowFactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meowFactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meowFactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
