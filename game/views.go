package game

import (
	"fmt"
	"time"

	"github.com/calvinmclean/terminal-tiller/farm"
	"github.com/calvinmclean/terminal-tiller/styles"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const readyToHarvestMsg = "Ready to harvest!"

func (g *game) statusBar() string {
	w := lipgloss.Width

	statusBarStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	moneyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1).
		Background(lipgloss.Color("#A550DF")).
		Align(lipgloss.Right)

	statusText := lipgloss.NewStyle().Inherit(statusBarStyle)

	statusKey := statusBarStyle.Render(g.farm.Name())
	money := moneyStyle.Render(fmt.Sprintf("%dg", g.farm.Money()))
	statusVal := statusText.Copy().
		Width(g.actualWidth - w(statusKey) - w(money)).
		Render("")

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		money,
	)

	return statusBarStyle.Width(g.actualWidth).Render(bar)
}

func (g *game) renderTable() string {
	tableData := g.farm.Data()

	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		// Width(actualWidth / 2).
		// Height(actualWidth).
		Rows(tableData...).
		StyleFunc(func(row, col int) lipgloss.Style {
			var bg lipgloss.TerminalColor = lipgloss.NoColor{}
			var fg lipgloss.TerminalColor = lipgloss.NoColor{}

			// background color of current selection
			c := coord{col, row - 1}
			if c == g.curCoord || g.isSelected(c) {
				bg = lipgloss.Color("241")

				current := g.farm.Get(c.row, c.col)
				if current == nil && // tile is NOT a plant
					g.selectedCropType.SeedCost()*g.numSelected() <= g.farm.Money() { // can afford to plant all tiles
					bg = styles.Green
				}
				if current != nil && current.ReadyToHarvest(g.farm.TimeScale()) {
					bg = styles.Yellow
				}
			}

			return lipgloss.NewStyle().
				Foreground(fg).
				Background(bg).
				Padding(0, 1).
				Align(lipgloss.Center, lipgloss.Center)
		}).
		Render()
}

func (g *game) selectedSeedView() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		styles.ListHeader("Current Seed"),
		styles.ListItem(fmt.Sprintf("%s %s", g.selectedCropType.Name(), g.selectedCropType.Representation())),
		styles.ListItem(cropCost(g.selectedCropType)),
		styles.ListItem(styles.AquaText(fmt.Sprintf("%v", g.selectedCropType.HarvestTime(g.farm.TimeScale())))),
	)
}

func cropCost(cropType farm.CropType) string {
	return fmt.Sprintf(
		"%s / %s",
		styles.RedText(fmt.Sprintf("-%dg", cropType.SeedCost())),
		styles.GreenText(fmt.Sprintf("+%dg", cropType.MarketPrice())),
	)
}

func (g *game) selectedCellView() string {
	return lipgloss.NewStyle().
		Width(len(readyToHarvestMsg) + 2).
		Render(lipgloss.JoinVertical(lipgloss.Left, g.selectedCellDetails()...))
}

func (g *game) selectedCellDetails() []string {
	header := styles.ListHeader("Selected Cell")

	crop := g.farm.Get(g.curCoord.row, g.curCoord.col)
	if crop == nil {
		return []string{
			header,
			styles.ListItem("Soil"),
		}
	}

	harvestDate := crop.PlantedAt.Add(crop.Type.HarvestTime(g.farm.TimeScale()))
	timeUntilHarvest := time.Since(harvestDate) * -1
	timeUntilHarvestDisplay := styles.YellowText(fmt.Sprintf("%v", timeUntilHarvest.Truncate(time.Millisecond)))
	if timeUntilHarvest < 0 {
		timeUntilHarvestDisplay = styles.GreenText(readyToHarvestMsg)
	}

	return []string{
		header,
		styles.ListItem(fmt.Sprintf("%s %s", crop.Type.Name(), crop.Type.Representation())),
		styles.ListItem(progressBar(crop, g.farm.TimeScale())),
		styles.ListItem(timeUntilHarvestDisplay),
	}
}

func progressBar(crop *farm.Crop, scale time.Duration) string {
	progress := progress.New(
		progress.WithGradient(string(styles.Yellow), string(styles.Green)),
		progress.WithWidth(len(readyToHarvestMsg)),
		progress.WithoutPercentage(),
	)

	harvestDate := crop.HarvestTime(scale)
	until := time.Until(harvestDate)
	percent := 1 - float64(until)/float64(crop.Type.HarvestTime(scale))

	return progress.ViewAs(percent)
}
