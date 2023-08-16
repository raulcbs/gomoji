/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/raulcbs/gomoji/pkg"
	"github.com/spf13/cobra"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape emojis from gitmoji.dev.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scrape called")

		emojis := pkg.GetMojis()

		for _, emoji := range emojis {
			fmt.Printf("\nIcon: %v\nCode: %v\nName: %v\n", emoji.Icon, emoji.Code, emoji.Name)
			fmt.Println("=========================================================")
		}
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}
