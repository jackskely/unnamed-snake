package main

import "github.com/jackskely/unnamed-snake/internal/game"

func main() {
	snakeGame := game.NewController("Unnamed Snake Game", 800, 800, 60)
	snakeGame.Start(800, 800, 10, 4)
}
