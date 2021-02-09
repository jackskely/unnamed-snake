package game

import (
	"log"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	log.Println(seed)
}

type game struct {
	width, height, axisCellNumber int32
	input                         chan direction
	food                          coord
	snake                         snake
}

// NewGame returns a new game instance
func NewGame(width, height, axisCellNumber int32) game {
	snk := newSnake(axisCellNumber)
	return game{
		width:          width,
		height:         height,
		axisCellNumber: axisCellNumber,
		food:           newFood(axisCellNumber, snk),
		input:          make(chan direction),
		snake:          snk,
	}
}

// Start starts the game instance
func (g *game) Start() {
	go g.gameLoop()
	g.drawLoop()
}

func (g *game) gameLoop() {
	go g.listenInput()
	<-g.input
	max := int(g.width)*int(g.height) - 1
	for range time.Tick(time.Second / 8) {
		if len(g.snake.body) == max {
			log.Fatalf("VICTOIRE: %+v", g)
		}
		if g.snake.body[0].pos.x == g.food.x && g.snake.body[0].pos.y == g.food.y {
			g.snake.eat()
			g.food = newFood(g.axisCellNumber, g.snake)
		}
		g.snake.move()
		if g.snake.body[0].pos.x >= g.axisCellNumber || g.snake.body[0].pos.y >= g.axisCellNumber || g.snake.body[0].pos.x < 0 || g.snake.body[0].pos.y < 0 {
			log.Fatalf("MUR: %+v", g)
		}
		for i := 1; i < len(g.snake.body); i++ {
			if g.snake.body[0].pos.x == g.snake.body[i].pos.x && g.snake.body[0].pos.y == g.snake.body[i].pos.y {

				log.Fatalf("SUICIDE: %+v", g)
			}
		}
	}
}

func (g *game) drawLoop() {
	rl.InitWindow(g.width, g.height, "Unnamed Snake Game")
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		g.drawFood()
		g.drawSnake()
		rl.DrawFPS(2, 2)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func captureInput(input chan<- direction) {
	for range time.Tick(time.Second / 16) {
		if rl.IsKeyDown(rl.KeyLeft) {
			input <- left
		} else if rl.IsKeyDown(rl.KeyDown) {
			input <- down
		} else if rl.IsKeyDown(rl.KeyRight) {
			input <- right
		} else if rl.IsKeyDown(rl.KeyUp) {
			input <- up
		}
	}
}

func (g *game) listenInput() {
	go captureInput(g.input)
	for {
		select {
		case d := <-g.input:
			g.snake.body[0].dir = d
		}
	}
}

func newFood(axisCellNumber int32, s snake) coord {
	rnx := rand.Int31n(axisCellNumber)
	rny := rand.Int31n(axisCellNumber)
	for _, cell := range s.body {
		if cell.pos.y == rnx && cell.pos.y == rny {
			return newFood(axisCellNumber, s)
		}
	}
	return coord{x: rnx, y: rny}
}

func (g game) drawSnake() {
	width := g.width / g.axisCellNumber
	height := g.height / g.axisCellNumber

	switch g.snake.body[0].dir {
	case left:
		rl.DrawRectangleGradientH(g.snake.body[0].pos.x*(width), g.snake.body[0].pos.y*(height), width, height, rl.DarkGreen, rl.Green)
	case down:
		rl.DrawRectangleGradientV(g.snake.body[0].pos.x*(width), g.snake.body[0].pos.y*(height), width, height, rl.Green, rl.DarkGreen)
	case up:
		rl.DrawRectangleGradientV(g.snake.body[0].pos.x*(width), g.snake.body[0].pos.y*(height), width, height, rl.DarkGreen, rl.Green)
	case right:
		rl.DrawRectangleGradientH(g.snake.body[0].pos.x*(width), g.snake.body[0].pos.y*(height), width, height, rl.Green, rl.DarkGreen)
	}

	for _, s := range g.snake.body[1:] {
		rl.DrawRectangle(s.pos.x*(width), s.pos.y*(height), width, height, rl.Green)
	}
}

func (g game) drawFood() {
	rl.DrawRectangle(g.food.x*(g.width/g.axisCellNumber), g.food.y*(g.height/g.axisCellNumber), g.width/g.axisCellNumber, g.height/g.axisCellNumber, rl.Red)
}

func (g *game) log() {
	for range time.Tick(time.Second) {
		log.Printf("%+v", g.snake.body)
		log.Printf("%+v,%+v,%+v,%+v",
			g.snake.body[0].pos.x >= g.axisCellNumber, g.snake.body[0].pos.y >= g.axisCellNumber, g.snake.body[0].pos.x < 0, g.snake.body[0].pos.y < 0)
	}
}
