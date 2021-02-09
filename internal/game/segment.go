package game

type segment struct {
	pos coord
	dir direction
}

func (c *segment) next() {
	switch c.dir {
	case left:
		c.pos.left(1)
	case down:
		c.pos.down(1)
	case up:
		c.pos.up(1)
	case right:
		c.pos.right(1)
	}
}

func (c *segment) before() {
	switch c.dir {
	case left:
		c.pos.right(1)
	case down:
		c.pos.up(1)
	case up:
		c.pos.down(1)
	case right:
		c.pos.left(1)
	}
}
