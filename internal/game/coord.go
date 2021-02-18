package game

type coord struct {
	x, y int32
}

func (c coord) equal(c1 coord) bool {
	return c.x == c1.x && c.y == c1.y
}

func (c coord) isWithinRect(x1, y1, x2, y2 int32) bool {
	return x1 <= c.x && c.x < x2 && y1 <= c.y && c.y < y2
}
