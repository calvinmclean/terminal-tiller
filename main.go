package main

import (
	"github.com/calvinmclean/terminal-tiller/game"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	g := game.New()
	p := tea.NewProgram(g, tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}
