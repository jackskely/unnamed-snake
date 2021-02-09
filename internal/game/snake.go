package game

import (
	"math/rand"
)

type snake struct {
	body []segment
}

func newSnake(axisCellNumber int32) snake {
	return snake{
		body: append(
			make([]segment, 0, axisCellNumber*axisCellNumber),
			segment{
				pos: coord{
					x: rand.Int31n(axisCellNumber),
					y: rand.Int31n(axisCellNumber),
				},
				dir: direction(rand.Intn(4)),
			},
		),
	}
}

func (s *snake) head() *segment {
	return &s.body[0]
}

func (s *snake) move() {
	s.body[0].next()
	for i := len(s.body) - 1; i > 0; i-- {
		s.body[i].next()
		s.body[i].dir = s.body[i-1].dir
	}
}

func (s *snake) eat() {
	last := s.body[len(s.body)-1]
	last.before()
	s.body = append(s.body, last)
}
