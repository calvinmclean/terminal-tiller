package game

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/calvinmclean/terminal-tiller/farm"
	"github.com/calvinmclean/terminal-tiller/styles"

	"github.com/charmbracelet/bubbles/list"
)

func newSeedSelectView(scale time.Duration) list.Model {
	items := []list.Item{}
	for cropType := range farm.CropTypes {
		items = append(items, seedListItem{cropType, scale})
	}

	slices.SortFunc[[]list.Item](items, func(a, b list.Item) int {
		return strings.Compare(a.FilterValue(), b.FilterValue())
	})

	plantList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	plantList.Title = "Seeds"

	return plantList
}

type seedListItem struct {
	farm.CropType
	scale time.Duration
}

func (i seedListItem) Title() string {
	return fmt.Sprintf("%s %s", i.CropType.Representation(), i.CropType.Name())
}

func (i seedListItem) Description() string {
	return fmt.Sprintf(
		"%s %s",
		cropCost(i.CropType),
		styles.AquaText(fmt.Sprintf("%v", i.CropType.HarvestTime(i.scale))),
	)
}

func (i seedListItem) FilterValue() string {
	return fmt.Sprintf("%03d", i.CropType.SeedCost())
}
