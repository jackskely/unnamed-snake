package game

import (
	"log"
	"time"
)

type world struct {
	width, height         int32
	cellWidth, cellHeight int32
	axisCellNumber        int32
	score                 int32
	food                  food
	snake                 snake
}

func newWorld(width, height, axisCellNumber int32) world {
	return world{
		width:          width,
		height:         height,
		cellWidth:      width / axisCellNumber,
		cellHeight:     height / axisCellNumber,
		axisCellNumber: axisCellNumber,
		food:           newRandomFood(axisCellNumber),
		snake:          newRandomSnake(axisCellNumber),
	}
}

func (w *world) starting(tps int32, c <-chan direction) {
	tick := time.NewTicker(time.Second / time.Duration(tps))
	<-c
	for range tick.C {
		w.snake.move()

		if failure(w) {
			break
		}

		if growthRule(w) {
			w.snake.eat()
			w.score += 3 * int32(len(w.snake.body))
			if success(w) {
				break
			}
			w.food = foodSpawnMechanic(w)
		}
	}

	tick.Stop()
}

func success(w *world) bool {
	return maxSizeRule(w)
}

func failure(w *world) bool {
	return outOfBoundsRule(w) || ouroborosRule(w)
}

func maxSizeRule(w *world) bool {
	if len(w.snake.body) == int(w.axisCellNumber)*int(w.axisCellNumber) {
		log.Println("DRAGON", w.snake)
		return true
	}
	return false
}

func outOfBoundsRule(w *world) bool {
	head := w.snake.body[0]
	if !head.isWithinRect(0, 0, w.axisCellNumber, w.axisCellNumber) {
		log.Println("WONDERWALL", head)
		return true
	}
	return false
}

func ouroborosRule(w *world) bool {
	head := w.snake.body[0]
	for _, s := range w.snake.body[1:] {
		if s.equal(head) {
			log.Println("OUROBOROS", w.snake)
			return true
		}
	}
	return false
}

func growthRule(w *world) bool {
	head := w.snake.body[0]
	return head.equal(coord(w.food))
}

func foodSpawnMechanic(w *world) food {
	f := newRandomFood(w.axisCellNumber)
	for _, s := range w.snake.body {
		if s.equal(coord(f)) {
			return foodSpawnMechanic(w)
		}
	}
	return f
}
