/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"cobra-kai/model"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// promptDavinciTwoCmd represents the promptDavinciTwo command
var promptDavinciTwoCmd = &cobra.Command{
	Use:   "promptDavinciTwo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		freshApiKey, _ := godotenv.Read(".env")
		prompt, _ := cmd.Flags().GetString("prompt")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"model":       "text-davinci-002",
			"prompt":      prompt,
			"max_tokens":  2000,
			"temperature": 0,
		})

		auth := "Bearer " + freshApiKey["OPENAI_API_KEY"]

		req, _ := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", auth)

		resp, _ := http.DefaultClient.Do(req)

		defer resp.Body.Close()

		textCompletion := new(model.TextCompletion)

		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &textCompletion)
		if strings.Contains(string(body), "You didn't provide an API key") {
			fmt.Println("You need to provide API key")
			return
		} else {
			fmt.Println(textCompletion.Choices[0].Text)
		}
	},
}

func init() {
	rootCmd.AddCommand(promptDavinciTwoCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptDavinciTwoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptDavinciTwoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
