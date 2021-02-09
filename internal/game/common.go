package game

type direction int

const (
	left direction = iota
	down
	right
	up
)

type event string

const (
	lost event = "You've lost"
	won        = "You've won"
)
