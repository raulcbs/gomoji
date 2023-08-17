package pkg

import (
	"github.com/raulcbs/gomoji/utils"
)

func ListEmojisAvaliable() ([]utils.Emoji, error) {
	var emojis []utils.Emoji

	utils.GetEmojisFromJSON(&emojis)

	return emojis, nil
}
