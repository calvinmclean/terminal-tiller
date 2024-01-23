package main

import (
	"flag"

	"github.com/calvinmclean/terminal-tiller/game"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	filename := flag.String("file", "", "save file to load from")
	farmName := flag.String("name", "", "use this flag to create new farm with the provided name. filename is required when creating a new farm")
	flag.Parse()

	g, err := game.New(*filename, *farmName)
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(g, tea.WithAltScreen())

	_, err = p.Run()
	if err != nil {
		panic(err)
	}
}
