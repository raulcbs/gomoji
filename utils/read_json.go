package utils

import (
	"encoding/json"
	"log"
	"os"
)

func GetEmojisFromJSON(emojis *[]Emoji) {
	bytValue, err := os.ReadFile("emojis.json")
	if err != nil {
		log.Default().Println("error to read the json with contents of emojis")
	}

	json.Unmarshal(bytValue, &emojis)
}
