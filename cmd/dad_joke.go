/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// dadJokeCmd represents the dadJoke command
var dadJokeCmd = &cobra.Command{
	Use:     "dadJoke",
	Aliases: []string{"dad-joke"},
	Short:   "Prints out a dad joke",
	Long:    `Funny joke`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := ReadConfig()
		if err != nil {
			fmt.Println("reading config", err)
			return
		}

		client := &http.Client{}
		request, err := http.NewRequest(
			http.MethodGet,
			cfg.DadJokeBaseURL.String(),
			nil,
		)
		if err != nil {
			fmt.Println("create error:", err)
			return
		}

		request.Header.Add("Accept", "application/json")
		request.Header.Add("User-Agent", "cobra cli application - github.com/dom-m17/talk-to-me")

		// Make the request
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println("get error:", err)
			return
		}

		// Read the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		// Unmarshal the response
		var joke DadJoke
		err = json.Unmarshal(body, &joke)
		if err != nil {
			fmt.Println("Unmarshal error:", err)
			return
		}

		// Print joke
		fmt.Println(joke.Joke)
	},
}

func init() {
	rootCmd.AddCommand(dadJokeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dadJokeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dadJokeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
