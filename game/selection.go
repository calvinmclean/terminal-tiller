package game

import "github.com/calvinmclean/terminal-tiller/farm"

func (g *game) stopSelecting() {
	g.selecting = false
	g.selectedCoord = coord{-1, -1}
}

func (g *game) handleSelection() {
	if g.selecting {
		return
	}

	g.selectedCoord = g.curCoord
	g.selecting = true
}

// doForCurrentSelection will modifRow each cell in the currentlRow-selected range
func (g *game) doForCurrentSelection(do func(int, int)) {
	startCol, startRow, endCol, endRow := g.getCurrentSelections()
	for row := startRow; row <= endRow; row++ {
		for col := startCol; col <= endCol; col++ {
			do(row, col)
		}
	}
}

// getCurrentSelections returns the current Col/Row and the last selected Col/Row in order with top-left first
func (g *game) getCurrentSelections() (int, int, int, int) {
	return g.getCurrentSelectionsAtCoord(g.curCoord)
}

func (g *game) getCurrentSelectionsAtCoord(curCoord coord) (int, int, int, int) {
	startCol, endCol := g.selectedCoord.col, curCoord.col
	startRow, endRow := g.selectedCoord.row, curCoord.row

	if startRow > endRow {
		startRow, endRow = endRow, startRow
	}
	if startCol > endCol {
		startCol, endCol = endCol, startCol
	}

	if startCol == -1 {
		startCol = endCol
	}
	if startRow == -1 {
		startRow = endRow
	}

	return startCol, startRow, endCol, endRow
}

func (g *game) isSelected(c coord) bool {
	if !g.selecting {
		return false
	}
	startCol, startRow, endCol, endRow := g.getCurrentSelections()
	return c.col >= startCol && c.col <= endCol && c.row >= startRow && c.row <= endRow
}

func (g *game) numSelected() int {
	startCol, startRow, endCol, endRow := g.getCurrentSelections()
	colSize := endCol - startCol + 1
	rowSize := endRow - startRow + 1

	return colSize * rowSize
}

func (g *game) getCurrentCell() *farm.Crop {
	return g.farm.Get(g.curCoord.row, g.curCoord.col)
}
