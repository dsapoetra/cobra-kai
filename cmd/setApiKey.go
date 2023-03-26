/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cobra-kai/utils"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
)

// setApiKeyCmd represents the setApiKey command
var setApiKeyCmd = &cobra.Command{
	Use:   "setApiKey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, _ := cmd.Flags().GetString("api-key")
		utils.WriteToDotEnv(apiKey)
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading environment variables file")
		}

		freshApiKey, _ := godotenv.Read(".env")

		fmt.Println("OPENAI_API_KEY set to: ", freshApiKey["OPENAI_API_KEY"])
	},
}

func init() {
	rootCmd.AddCommand(setApiKeyCmd)
	setApiKeyCmd.PersistentFlags().StringP("api-key", "a", "", "use prompt for openai prompt")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setApiKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setApiKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
