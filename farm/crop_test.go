package farm

import (
	"fmt"
	"testing"
)

// This isn't a real test but is useful for checking crop details and balancing
func TestDisplayCropStats(t *testing.T) {
	for cropType := range CropTypes {
		profit := cropType.MarketPrice() - cropType.SeedCost()
		ppd := float64(profit) / float64(cropType.HarvestTime(1))
		fmt.Printf("%-12s: -%3dg / +%3dg = +%3d  ppd: %.3f\n", cropType.Name(), cropType.SeedCost(), cropType.MarketPrice(), profit, ppd)
	}
}
