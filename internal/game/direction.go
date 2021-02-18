package game

import "math/rand"

type direction int

const (
	up direction = iota
	left
	down
	right
)

func randomDirection() direction {
	return direction(rand.Int31n(4))
}
