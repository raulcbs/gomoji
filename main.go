package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

func doCommit() {
	emojis := []string{
		"✨ :sparkles: Nueva característica",
		":bug: Corrección de errores",
		":fire: Eliminación de código",
		":hammer: Cambios en la estructura",
		":art: Mejora en el formato/código",
		// Agrega más emoticonos y descripciones según tus necesidades
	}

	color.Yellow("Selecciona un emoticono:")
	for i, emoji := range emojis {
		color.Cyan("%d. %s", i+1, emoji)
	}

	reader := bufio.NewReader(os.Stdin)
	color.Yellow("Ingresa el número del emoticono deseado: ")
	selection, _ := reader.ReadString('\n')

	var selectedEmoji string
	selectedIndex := -1
	color.Yellow("Escribe el mensaje del commit: ")
	commitMessage, _ := reader.ReadString('\n')

	fmt.Sscanf(selection, "%d", &selectedIndex)
	if selectedIndex >= 1 && selectedIndex <= len(emojis) {
		selectedEmoji = emojis[selectedIndex-1]
		// cmd := exec.Command("git", "commit", "-m", selectedEmoji+" "+commitMessage)
		cmd := exec.Command("ls")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			color.Red("Error al ejecutar git commit:", err)
		} else {
			color.Green("Commit exitoso: %s %s", selectedEmoji, strings.TrimSpace(commitMessage))
		}
	} else {
		color.Red("Opción no válida. Selecciona un número del 1 al %d", len(emojis))
	}
}

type moji struct {
	emoji string
	code  string
	name  string
}

func main() {

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
