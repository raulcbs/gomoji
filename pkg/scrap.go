package pkg

import (
	"github.com/gocolly/colly"
)

type Emoji struct {
	Icon string
	Code string
	Name string
}

func GetMojis() []Emoji {
	var emojis []Emoji

	c := colly.NewCollector(
		colly.AllowedDomains("gitmoji.dev"),
	)

	c.OnHTML("article.styles_emoji__nVHNW", func(e *colly.HTMLElement) {
		singleEmoji := Emoji{}
		singleEmoji.Icon = e.ChildAttr("button.gitmoji-clipboard-emoji", "data-clipboard-text")
		singleEmoji.Code = e.ChildAttr("button.gitmoji-clipboard-code", "data-clipboard-text")
		singleEmoji.Name = e.ChildText("p")

		emojis = append(emojis, singleEmoji)
	})

	c.Visit("https://gitmoji.dev/")

	return emojis
}
