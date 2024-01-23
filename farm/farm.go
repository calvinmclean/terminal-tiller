package farm

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const startingMoney = 5000

type Farm struct {
	name      string
	w, h      int
	timeScale time.Duration
	field     [][]*Crop
	money     int
}

func New(name string, w, h int, scale time.Duration) *Farm {
	field := make([][]*Crop, h)
	for row := range field {
		field[row] = make([]*Crop, w)
	}

	return &Farm{
		name,
		w, h,
		scale,
		field,
		startingMoney,
	}
}

func Load(data []byte) (*Farm, error) {
	decrypted, err := decrypt(data)
	if err != nil {
		return nil, fmt.Errorf("error decrypting: %w", err)
	}

	var f encodeableFarm
	err = json.Unmarshal(decrypted, &f)
	if err != nil {
		return nil, fmt.Errorf("error parsing farm data: %w", err)
	}

	return &Farm{
		f.Name,
		f.W, f.H,
		f.TimeScale,
		f.Field,
		f.Money,
	}, nil
}

func (f *Farm) Data() [][]string {
	tableData := make([][]string, f.h)
	for row := range f.field {
		tableData[row] = make([]string, f.w)
		for col, crop := range f.field[row] {
			switch {
			case crop != nil:
				tableData[row][col] = crop.String(f.timeScale)
			default:
				tableData[row][col] = " "
			}
		}
	}

	return tableData
}

func (f *Farm) Name() string {
	return f.name
}

func (f *Farm) TimeScale() time.Duration {
	return f.timeScale
}

func (f *Farm) Height() int {
	return f.h
}

func (f *Farm) Width() int {
	return f.w
}

func (f *Farm) Money() int {
	return f.money
}

func (f *Farm) Get(row, col int) *Crop {
	return f.field[row][col]
}

func (f *Farm) Plant(cropType CropType, row, col int) error {
	cur := f.Get(row, col)
	if cur != nil {
		return errors.New("crop here already")
	}

	if f.money < cropType.SeedCost() {
		return errors.New("not enough money")
	}

	f.money -= cropType.SeedCost()
	f.field[row][col] = NewCrop(cropType)

	return nil
}

func (f *Farm) Harvest(row, col int) error {
	cur := f.Get(row, col)
	if cur == nil {
		return errors.New("no crop here")
	}
	if !cur.ReadyToHarvest(f.TimeScale()) {
		return errors.New("not ready to harvest")
	}

	f.money += cur.Type.MarketPrice()
	f.field[row][col] = nil

	return nil
}
