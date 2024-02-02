package main

import (
	"flag"
	"fmt"

	"github.com/calvinmclean/terminal-tiller/game"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	filename := flag.String("file", "", "save file to load from")
	farmName := flag.String("name", "", "use this flag to create new farm with the provided name. filename is required when creating a new farm")
	status := flag.Bool("status", false, "print out the game status")
	flag.Parse()

	g, f, err := game.New(*filename, *farmName)
	if err != nil {
		panic(err)
	}

	if *status {
		fmt.Println(f.Status())
		return
	}

	p := tea.NewProgram(g, tea.WithAltScreen())

	_, err = p.Run()
	if err != nil {
		panic(err)
	}
}
