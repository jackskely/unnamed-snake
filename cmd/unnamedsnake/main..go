package main

import (
	"github.com/jackskely/unnamed-snake/internal/game"
)

func main() {
	g := game.NewGame(800, 800, 16)
	g.Start()
}
