/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/raulcbs/gomoji/utils"
	"github.com/spf13/cobra"
)

var Description string
var Code string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search emoji you want",
	Long: `
		With this subcommand you can search the emoji you want to use
		in your commit. You can search by code ' -c :memo: ' or by
		description ' -d "update" '. 
		You don't need the write all the description or code to search.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if Description != "" && Code != "" {
			fmt.Println("You must use only my flag")
			return
		}

		if Description != "" {
			fmt.Println("")
			var emojis []utils.Emoji

			utils.GetEmojisFromJSON(&emojis)

			for _, emoji := range emojis {
				if strings.Contains(emoji.Name, Description) {
					if count, _ := fmt.Printf("%v -- %v -- %v\n", emoji.Icon, emoji.Code, emoji.Name); count == 0 {
						fmt.Printf("Don't have any emoji with : %v \n", Description)
					}
				}
			}
			fmt.Println("")
			return
		}

		if Code != "" {
			fmt.Println("")
			var emojis []utils.Emoji

			utils.GetEmojisFromJSON(&emojis)

			for _, emoji := range emojis {
				if strings.Contains(emoji.Code, Code) {
					if count, _ := fmt.Printf("%v -- %v -- %v\n", emoji.Icon, emoji.Code, emoji.Name); count == 0 {
						fmt.Printf("Don't have any emoji with : %v \n", Code)
					}
				}
			}
			fmt.Println("")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&Description, "description", "d", "", "Search by emoji description")
	searchCmd.Flags().StringVarP(&Code, "code", "c", "", "Search by emoji code")
}
