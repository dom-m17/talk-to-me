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

type Quote struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := ReadConfig()
		if err != nil {
			fmt.Println("reading config:", err)
		}

		client := &http.Client{}
		request, err := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf("%s%s", cfg.ApiNinjaBaseURL.String(), "quotes"),
			nil,
		)
		if err != nil {
			fmt.Println("create error:", err)
			return
		}

		request.Header.Add("Accept", "application/json")
		request.Header.Add("X-Api-Key", cfg.ApiNinjaAPIKey)
		request.Header.Add("User-Agent", "cobra cli application - github.com/dom-m17/talk-to-me")

		resp, err := client.Do(request)
		if err != nil {
			fmt.Println("get error:", err)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		var quote []Quote
		err = json.Unmarshal(body, &quote)
		if err != nil {
			fmt.Println("Unmarshal error:", err)
			return
		}

		fmt.Println(quote[0].Quote)
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
