package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type moji struct {
	emoji string
	code  string
	name  string
}

func getMojis() {
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
}
