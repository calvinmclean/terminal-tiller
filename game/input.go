package game

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (g *game) handleInput(msg tea.KeyMsg) tea.Cmd {
	if g.showSeedSelect {
		return g.handleListInput(msg)
	}

	switch msg.String() {
	case "ctrl+c", "q":
		return g.saveAndQuit
	case "up", "k",
		"down", "j",
		"left", "h",
		"right", "l":
		g.move(msg)
	case "esc":
		g.stopSelecting()
	case "s":
		g.showSeedSelect = true
	case "f":
		g.interact()
	case "enter", " ":
		g.handleSelection()
	}
	return nil
}

func (g *game) handleListInput(msg tea.KeyMsg) tea.Cmd {
	switch msg.String() {
	case "s", "q":
		g.showSeedSelect = false
	case "enter", " ":
		g.selectedCropType = g.seedSelect.SelectedItem().(seedListItem).CropType
		g.showSeedSelect = false
	}
	return nil
}

func (g *game) move(msg tea.KeyMsg) {
	newCoord := coord{g.curCoord.col, g.curCoord.row}
	switch msg.String() {
	case "up", "k":
		newCoord.up()
	case "down", "j":
		newCoord.down(g.farm.Height())
	case "left", "h":
		newCoord.left()
	case "right", "l":
		newCoord.right(g.farm.Width())
	}

	if !g.selecting {
		g.curCoord = newCoord
		return
	}

	// don't allow overlap of existing plants or outside of harvestable area
	selected := g.farm.Get(g.selectedCoord.row, g.selectedCoord.col)
	startCol, startRow, endCol, endRow := g.getCurrentSelectionsAtCoord(newCoord)
	for row := startRow; row <= endRow; row++ {
		for col := startCol; col <= endCol; col++ {
			current := g.farm.Get(row, col)

			// Only select crops of the same type
			if current != nil && selected != nil && current.Type != selected.Type {
				return
			}

			// All selections must be crop or no crop
			if (current == nil) != (selected == nil) {
				return
			}
		}
	}
	g.curCoord = newCoord
}

func (g *game) interact() {
	crop := g.getCurrentCell()
	switch crop {
	case nil:
		g.plant()
	default:
		g.harvest()
	}
}

func (g *game) plant() {
	// can't plant if there's not enough money
	if g.selectedCropType.SeedCost()*g.numSelected() > g.farm.Money() {
		return
	}

	g.doForCurrentSelection(func(row, col int) {
		g.farm.Plant(g.selectedCropType, row, col)
	})
	g.stopSelecting()
}

func (g *game) harvest() {
	g.doForCurrentSelection(func(row, col int) {
		g.farm.Harvest(row, col)
	})
	g.stopSelecting()
}
