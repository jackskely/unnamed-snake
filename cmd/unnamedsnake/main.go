package main

import "github.com/jackskely/unnamed-snake/internal/game"

func main() {
	snakeGame := game.NewController()
	snakeGame.Start(800, 800, 10, "Unnamed Snake Game", 60, 4)
}
