/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

// struct of Moji we have
type moji struct {
	emoji string
	code  string
	name  string
}

// scrapmojisCmd represents the scrapmojis command
var scrapmojisCmd = &cobra.Command{
	Use:   "scrapmojis",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var emojis []moji

		c := colly.NewCollector(
			colly.AllowedDomains("gitmoji.dev"),
		)

		c.OnHTML("article.styles_emoji__nVHNW", func(e *colly.HTMLElement) {
			singleEmoji := moji{}

			singleEmoji.emoji = e.ChildAttr("button.gitmoji-clipboard-emoji", "data-clipboard-text")
			singleEmoji.code = e.ChildAttr("button.gitmoji-clipboard-code", "data-clipboard-text")
			singleEmoji.name = e.ChildText("p")

			fmt.Printf("Emoji: %v\n Code: %v\n Name: %v\n", singleEmoji.emoji, singleEmoji.code, singleEmoji.name)

			emojis = append(emojis, singleEmoji)
		})

		c.Visit("https://gitmoji.dev/")
	},
}

func init() {
	rootCmd.AddCommand(scrapmojisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapmojisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapmojisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
