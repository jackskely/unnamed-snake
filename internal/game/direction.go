package game

import "math/rand"

type control int

const (
	up control = iota
	left
	down
	right
)

func randomDirection() control {
	return control(rand.Int31n(4))
}
