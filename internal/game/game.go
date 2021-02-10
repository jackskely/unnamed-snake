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
	width, height         int32
	cellWidth, cellHeight int32
	axisCellNumber        int32
	food                  coord
	input                 chan direction
	draw                  chan event
	snake                 snake
}

// NewGame returns a new game instance
func NewGame(width, height, axisCellNumber int32) game {
	snk := newSnake(axisCellNumber)
	return game{
		width:          width,
		height:         height,
		cellWidth:      width / axisCellNumber,
		cellHeight:     height / axisCellNumber,
		axisCellNumber: axisCellNumber,
		food:           newFood(axisCellNumber, &snk),
		input:          make(chan direction),
		draw:           make(chan event),
		snake:          snk,
	}
}

// Start starts the game instance
func (g *game) Start() {
	rl.InitWindow(g.width, g.height, "Unnamed Snake Game")
	rl.SetTargetFPS(60)

	go g.gameLoop()
	g.drawLoop()

	rl.CloseWindow()
}

func (g *game) gameLoop() {
	go g.listenInput()
	<-g.input

	max := int(g.width) * int(g.height)

	for range time.Tick(time.Second / 8) {
		if len(g.snake.body) == max {
			log.Print("Evolution: dragon")
			g.draw <- won
			return
		}

		if g.snake.head().x == g.food.x && g.snake.head().y == g.food.y {
			g.snake.eat()
			g.food = newFood(g.axisCellNumber, &g.snake)
		}

		g.snake.move()

		if g.snake.head().x >= g.axisCellNumber || g.snake.head().y >= g.axisCellNumber || g.snake.head().x < 0 || g.snake.head().y < 0 {
			log.Print("Death: wonderwall")
			g.draw <- lost
			return
		}

		for i := 1; i < len(g.snake.body); i++ {
			if g.snake.head().x == g.snake.body[i].x && g.snake.head().y == g.snake.body[i].y {
				log.Print("Death: ouroboros")
				g.draw <- lost
				return
			}
		}
	}
}

func (g *game) drawLoop() {

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		g.drawFood()
		g.drawSnake()

		select {
		case e := <-g.draw:
			rl.DrawText(string(e), 190, 200, 20, rl.Black)
			return
		default:
		}

		rl.DrawFPS(2, 2)

		rl.EndDrawing()
	}

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
			g.snake.dir = d
		}
	}
}

func newFood(axisCellNumber int32, s *snake) coord {
	rnx := rand.Int31n(axisCellNumber)
	rny := rand.Int31n(axisCellNumber)
	for _, seg := range s.body {
		if seg.y == rnx && seg.y == rny {
			return newFood(axisCellNumber, s)
		}
	}
	return coord{x: rnx, y: rny}
}

func (g game) drawSnake() {
	head := g.snake.head()
	switch g.snake.dir {
	case left:
		rl.DrawRectangleGradientH(head.x*g.cellWidth, head.y*g.cellHeight, g.cellWidth, g.cellHeight, rl.DarkGreen, rl.Green)
	case down:
		rl.DrawRectangleGradientV(head.x*g.cellWidth, head.y*g.cellHeight, g.cellWidth, g.cellHeight, rl.Green, rl.DarkGreen)
	case up:
		rl.DrawRectangleGradientV(head.x*g.cellWidth, head.y*g.cellHeight, g.cellWidth, g.cellHeight, rl.DarkGreen, rl.Green)
	case right:
		rl.DrawRectangleGradientH(head.x*g.cellWidth, head.y*g.cellHeight, g.cellWidth, g.cellHeight, rl.Green, rl.DarkGreen)
	}
	for _, seg := range g.snake.body[1:] {
		rl.DrawRectangle(seg.x*g.cellWidth, seg.y*g.cellHeight, g.cellWidth, g.cellHeight, rl.Green)
	}
}

func (g game) drawFood() {
	rl.DrawRectangle(g.food.x*g.cellWidth, g.food.y*g.cellHeight, g.cellWidth, g.cellHeight, rl.Red)
}
