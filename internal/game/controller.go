package game

import (
	"log"
	"math/rand"
	"time"
)

type controller struct {
	view  view
	input inputHandler
}

func NewController(title string, width, height int32, targetFPS int32) controller {
	r := newRaylib(width, height, title, targetFPS)
	return controller{
		view:  &r,
		input: &r,
	}
}

func (c *controller) Start(width, height, axisCellNumber int32, tps int32) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	log.Printf("SEED: %d", seed)

	c.view.initView()

	w := newWorld(width, height, axisCellNumber)

	directionC := c.input.handleInput()
	worldC := w.start(tps, directionC)
	c.view.drawing(worldC)

	c.view.closeView()
}
