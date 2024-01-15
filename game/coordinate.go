package game

type coord struct {
	col, row int
}

func (c *coord) up() {
	if c.row > 0 {
		c.row--
	}
}

func (c *coord) down(h int) {
	if c.row < h-1 {
		c.row++
	}
}

func (c *coord) left() {
	if c.col > 0 {
		c.col--
	}
}

func (c *coord) right(w int) {
	if c.col < w-1 {
		c.col++
	}
}
