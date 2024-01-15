package styles

import "github.com/charmbracelet/lipgloss"

const (
	Yellow = lipgloss.Color("#F0E68C")
	Green  = lipgloss.Color("#73F59F")
)

var (
	GreenText  = lipgloss.NewStyle().Foreground(lipgloss.Color(Green)).Render
	RedText    = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render
	AquaText   = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Render
	YellowText = lipgloss.NewStyle().Foreground(lipgloss.Color(Yellow)).Render

	Subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	ListHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(Subtle).
			MarginRight(2).
			MarginLeft(2).
			Render

	ListItem = lipgloss.NewStyle().PaddingLeft(2).PaddingRight(2).Render
)
