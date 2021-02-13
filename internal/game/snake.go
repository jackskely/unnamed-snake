package game

import (
	"math/rand"
)

type snake struct {
	dir  direction
	body []coord
}

func newSnake(axisCellNumber int32) snake {
	return snake{
		body: append(
			make([]coord, 0, axisCellNumber*axisCellNumber),
			coord{
				x: rand.Int31n(axisCellNumber),
				y: rand.Int31n(axisCellNumber),
			},
		),
		dir: direction(rand.Intn(4)),
	}
}

func (s *snake) head() *coord {
	return &s.body[0]
}

func (s *snake) move() {
	head := s.body[0]
	switch s.dir {
	case left:
		head.x--
		prependAndPop(s, head)
	case down:
		head.y++
		prependAndPop(s, head)
	case up:
		head.y--
		prependAndPop(s, head)
	case right:
		head.x++
		prependAndPop(s, head)
	}
}

func (s *snake) eat() {
	s.body = append(s.body, s.body[len(s.body)-1])
}

func prependAndPop(s *snake, c coord) {
	copy(s.body[1:], s.body)
	s.body[0] = c
}
