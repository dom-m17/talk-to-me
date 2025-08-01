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

// factCmd represents the fact command
var factCmd = &cobra.Command{
	Use: "fact",
	// Short: "",
	Long: `Generates a random fact.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := ReadConfig()
		if err != nil {
			fmt.Println("reading config", err)
			return
		}

		client := &http.Client{}
		request, err := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf("%s%s", cfg.ApiNinjaBaseURL.String(), "facts"),
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

		var fact []Fact
		err = json.Unmarshal(body, &fact)
		if err != nil {
			fmt.Println("Unmarshal error:", err)
			return
		}

		fmt.Println(fact[0].Fact)
	},
}

func init() {
	rootCmd.AddCommand(factCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
