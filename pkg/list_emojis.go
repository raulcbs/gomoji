package pkg

import (
	"encoding/json"
	"errors"
	"os"
)

func ListEmojisAvaliable() ([]Emoji, error) {
	var emojis []Emoji

	bytValue, err := os.ReadFile("emojis.json")
	if err != nil {
		return []Emoji{}, errors.New("error to read the json with contents of emojis")
	}

	json.Unmarshal(bytValue, &emojis)

	return emojis, nil
}
