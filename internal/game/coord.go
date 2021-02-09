package game

type coord struct {
	x, y int32
}

func (c *coord) left(n int32) {
	c.x -= n
}
func (c *coord) down(n int32) {
	c.y += n
}
func (c *coord) up(n int32) {
	c.y -= n
}
func (c *coord) right(n int32) {
	c.x += n
}
