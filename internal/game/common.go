package game

type direction int

const (
	up direction = iota
	left
	down
	right
)

type event string

const (
	lost event = "You've lost"
	won  event = "You've won"
)
