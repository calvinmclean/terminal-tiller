package game

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/calvinmclean/terminal-tiller/farm"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	DEFAULT_WIDTH    = 7
	DEFAULT_HEIGHT   = 7
	DEFAULT_SCALE    = time.Minute
	DEFAULT_FILENAME = "my_farm.data"

	helpStr = `h/j/k/l or ←↓↑→ to move
enter or space to start a selection
c to cancel selection.
s to select seeds.
f to plant/harvest.
q to quit.
`
)

type game struct {
	farm *farm.Farm

	actualWidth, actualHeight int

	curCoord         coord
	selectedCropType farm.CropType

	selectedCoord coord
	selecting     bool

	showSeedSelect bool
	seedSelect     list.Model

	filename string
}

func New(filename, farmName string) (tea.Model, error) {
	if filename == "" {
		saveFiles, err := findSaveFiles()
		if err != nil {
			panic("error finding save files " + err.Error())
		}

		if len(saveFiles) > 0 {
			filename = saveFiles[0]
		}
	}

	dir, err := terminalTillerDir()
	if err != nil {
		return nil, fmt.Errorf("error determining save file directory: %w", err)
	}

	var f *farm.Farm
	switch {
	case filename == "" && farmName != "":
		return nil, fmt.Errorf("cannot create new farm without filename")
	case filename == "" && farmName == "": // create new farm with default name
		filename = filepath.Join(dir, DEFAULT_FILENAME)
		f = farm.New("My Farm", DEFAULT_WIDTH, DEFAULT_HEIGHT, DEFAULT_SCALE)
	case filename != "" && farmName == "": // load existing
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("error opening file: %w", err)
		}

		f, err = farm.Load(data)
		if err != nil {
			return nil, fmt.Errorf("error loading from save file: %w", err)
		}
	case filename != "" && farmName != "": // new with name
		f = farm.New(farmName, DEFAULT_WIDTH, DEFAULT_HEIGHT, DEFAULT_SCALE)
	}

	return &game{
		farm:             f,
		curCoord:         coord{},
		selectedCoord:    coord{-1, -1},
		seedSelect:       newSeedSelectView(f.TimeScale()),
		selectedCropType: farm.Lettuce,
		filename:         filename,
	}, nil
}

func (g *game) Init() tea.Cmd {
	return tick()
}

func (g *game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if g.showSeedSelect {
		g.seedSelect, cmd = g.seedSelect.Update(msg)
	}

	switch msg := msg.(type) {
	case tickMsg:
		cmd = tea.Batch(tick(), cmd)
	case tea.WindowSizeMsg:
		g.actualWidth = msg.Width
		g.actualHeight = msg.Height

		g.seedSelect.SetSize(g.actualWidth, g.actualHeight)
	case tea.KeyMsg:
		cmd = tea.Batch(g.handleInput(msg), cmd)
	}

	return g, cmd
}

func (g *game) View() string {
	lines := []string{}
	width := g.actualWidth

	seedSelectViewPort := ""
	if g.showSeedSelect {
		vpWidth := g.actualWidth / 4
		vp := viewport.New(vpWidth, g.actualHeight-4)
		vp.Style = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Margin(2, 0).
			PaddingRight(2)

		vp.SetContent(g.seedSelect.View())

		seedSelectViewPort = vp.View()
		width -= lipgloss.Width(seedSelectViewPort)
	}

	topView := lipgloss.PlaceHorizontal(
		width,
		lipgloss.Center,
		lipgloss.NewStyle().
			Margin(2, 2, 0).
			Render(
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					g.selectedSeedView(),
					g.selectedCellView(),
				),
			),
	)
	lines = append(lines, topView)

	table := lipgloss.PlaceHorizontal(
		width,
		lipgloss.Center,
		lipgloss.NewStyle().
			Margin(2, 2).
			Align(lipgloss.Center, lipgloss.Center).
			Render(g.renderTable()),
	)
	lines = append(lines, table)

	help := lipgloss.PlaceHorizontal(
		width,
		lipgloss.Center,
		lipgloss.NewStyle().
			Margin(2).
			Align(lipgloss.Center, lipgloss.Center).
			Render(helpStr),
	)
	lines = append(lines, help)

	result := lipgloss.JoinVertical(lipgloss.Top, lines...)
	if g.showSeedSelect {
		seedSelectViewPort = lipgloss.PlaceHorizontal(lipgloss.Width(seedSelectViewPort), lipgloss.Left, seedSelectViewPort)
		result = lipgloss.JoinHorizontal(lipgloss.Top, seedSelectViewPort, result)
	}

	statusBar := g.statusBar()
	statusBar = lipgloss.PlaceVertical(
		g.actualHeight-lipgloss.Height(result),
		lipgloss.Bottom, statusBar,
	)

	result = lipgloss.JoinVertical(lipgloss.Top, result, statusBar)

	return result
}
