package game

import "math/rand"

type snake struct {
	dir  control
	body []coord
}

func newSnake(x, y int32, dir control, totalCell int) snake {
	return snake{
		dir: dir,
		body: append(
			make([]coord, 0, totalCell),
			coord{x: x, y: y},
		),
	}
}

func newRandomSnake(axisCellNumber int32) snake {
	return newSnake(
		rand.Int31n(axisCellNumber),
		rand.Int31n(axisCellNumber),
		randomDirection(),
		int(axisCellNumber)*int(axisCellNumber),
	)
}

func (s *snake) move() {
	head := s.body[0]
	switch s.dir {
	case up:
		head.y--
		prependAndPop(s, head)
	case left:
		head.x--
		prependAndPop(s, head)
	case down:
		head.y++
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
