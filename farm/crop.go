package farm

import "time"

const (
	Broccoli CropType = iota
	Carrot
	Corn
	Cucumber
	Eggplant
	Garlic
	Ginger
	Lettuce
	Onion
	Pea
	Potato
	Strawberry
	Tomato
	Watermelon
	Yam
)

var (
	CropTypes = map[CropType]CropTypeDetails{
		Broccoli: {
			Name:           "Broccoli",
			Representation: "🥦",
			HarvestTime:    10,
			SeedCost:       80,
			MarketPrice:    150,
		},
		Carrot: {
			Name:           "Carrot",
			Representation: "🥕",
			HarvestTime:    4,
			SeedCost:       20,
			MarketPrice:    35,
		},
		Corn: {
			Name:           "Corn",
			Representation: "🌽",
			HarvestTime:    6,
			SeedCost:       30,
			MarketPrice:    55,
		},
		Cucumber: {
			Name:           "Cucumber",
			Representation: "🥒",
			HarvestTime:    5,
			SeedCost:       150,
			MarketPrice:    275,
		},
		Eggplant: {
			Name:           "Eggplant",
			Representation: "🍆",
			HarvestTime:    5,
			SeedCost:       125,
			MarketPrice:    200,
		},
		Garlic: {
			Name:           "Garlic",
			Representation: "🧄",
			HarvestTime:    15,
			SeedCost:       100,
			MarketPrice:    200,
		},
		Ginger: {
			Name:           "Ginger",
			Representation: "🫚 ",
			HarvestTime:    11,
			SeedCost:       90,
			MarketPrice:    175,
		},
		Lettuce: {
			Name:           "Lettuce",
			Representation: "🥬",
			HarvestTime:    3,
			SeedCost:       9,
			MarketPrice:    15,
		},
		Onion: {
			Name:           "Onion",
			Representation: "🧅",
			HarvestTime:    12,
			SeedCost:       80,
			MarketPrice:    155,
		},
		Pea: {
			Name:           "Pea",
			Representation: "🫛 ",
			HarvestTime:    3,
			SeedCost:       25,
			MarketPrice:    40,
		},
		Potato: {
			Name:           "Potato",
			Representation: "🥔",
			HarvestTime:    10,
			SeedCost:       50,
			MarketPrice:    80,
		},
		Strawberry: {
			Name:           "Strawberry",
			Representation: "🍓",
			HarvestTime:    6,
			SeedCost:       40,
			MarketPrice:    65,
		},
		Tomato: {
			Name:           "Tomato",
			Representation: "🍅",
			HarvestTime:    5,
			SeedCost:       70,
			MarketPrice:    130,
		},
		Watermelon: {
			Name:           "Watermelon",
			Representation: "🍉",
			HarvestTime:    8,
			SeedCost:       60,
			MarketPrice:    100,
		},
		Yam: {
			Name:           "Yam",
			Representation: "🍠",
			HarvestTime:    10,
			SeedCost:       75,
			MarketPrice:    160,
		},
	}
)

type CropType int

func (c CropType) Name() string {
	return CropTypes[c].Name
}

func (c CropType) Representation() string {
	return CropTypes[c].Representation
}

func (c CropType) HarvestTime(scale time.Duration) time.Duration {
	return time.Duration(CropTypes[c].HarvestTime) * scale
}

func (c CropType) SeedCost() int {
	return CropTypes[c].SeedCost
}

func (c CropType) MarketPrice() int {
	return CropTypes[c].MarketPrice
}

type CropTypeDetails struct {
	Name           string
	Representation string
	HarvestTime    int
	SeedCost       int
	MarketPrice    int
}

type Crop struct {
	Type      CropType
	PlantedAt time.Time
}

func NewCrop(cropType CropType) *Crop {
	_, ok := CropTypes[cropType]
	if !ok {
		panic("invalid crop type")
	}

	return &Crop{
		Type:      cropType,
		PlantedAt: time.Now(),
	}
}

func (c *Crop) String(scale time.Duration) string {
	if !c.ReadyToHarvest(scale) {
		return "🌱"
	}
	return c.Type.Representation()
}

func (c *Crop) ReadyToHarvest(scale time.Duration) bool {
	return time.Now().After(c.HarvestTime(scale))
}

func (c *Crop) HarvestTime(scale time.Duration) time.Time {
	return c.PlantedAt.Add(c.Type.HarvestTime(scale))
}
