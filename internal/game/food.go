package game

import "math/rand"

type food coord

func newFood(x, y int32) food {
	return food{x: x, y: y}
}

func newRandomFood(axisCellNumber int32) food {
	return newFood(
		rand.Int31n(axisCellNumber),
		rand.Int31n(axisCellNumber),
	)
}
