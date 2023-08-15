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
	Short: "Scrape emojis from gitmoji.dev and enrich your commits.",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
