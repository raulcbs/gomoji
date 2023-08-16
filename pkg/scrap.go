package pkg

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

type Emoji struct {
	Id   int    `json:"id"`
	Icon string `json:"icon"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func GetMojis() []Emoji {
	var emojis []Emoji
	var mu sync.Mutex // Mutex to ensure safe concurrent access

	c := colly.NewCollector(
		colly.AllowedDomains("gitmoji.dev"),
	)

	counter := 1 // Counter for generating unique IDs

	c.OnHTML("article.styles_emoji__nVHNW", func(e *colly.HTMLElement) {
		singleEmoji := Emoji{}
		singleEmoji.Id = counter // Assign unique ID
		counter++                // Increment counter

		singleEmoji.Icon = e.ChildAttr("button.gitmoji-clipboard-emoji", "data-clipboard-text")
		singleEmoji.Code = e.ChildAttr("button.gitmoji-clipboard-code", "data-clipboard-text")
		singleEmoji.Name = e.ChildText("p")

		mu.Lock()
		emojis = append(emojis, singleEmoji)
		mu.Unlock()
	})

	c.Visit("https://gitmoji.dev/")

	file, err := json.MarshalIndent(emojis, "", " ")
	if err != nil {
		log.Print("Error parsing struct to JSON:", err)
		return nil
	}

	err = os.WriteFile("emojis.json", file, 0644)
	if err != nil {
		log.Print("Error saving data to JSON:", err)
		return nil
	}

	return emojis
}
