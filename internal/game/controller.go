package game

import (
	"log"
	"math/rand"
	"time"
)

type controller struct {
	view     view
	input    inputHandler
	controlC chan direction
}

func NewController() controller {
	r := raylib{}
	return controller{
		view:     &r,
		input:    &r,
		controlC: make(chan direction),
	}
}

func (c *controller) Start(width, height, axisCellNumber int32, title string, targetFPS, tps int32) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	log.Printf("SEED: %d", seed)

	c.view.initView(width, height, title, targetFPS)

	go c.input.capturingInput(c.controlC)
	w := newWorld(width, height, axisCellNumber)
	go c.input.handlingInput(c.controlC, &w)
	go w.starting(tps, c.controlC)
	log.Printf("%+v", w)
	c.view.drawing(&w)

	c.view.closeView()
}
