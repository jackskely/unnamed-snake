package game

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type raylib struct {
	targetFPS int32
}

func (r *raylib) initView(width, height int32, title string, targetFPS int32) {
	rl.SetTraceLog(rl.LogNone)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(targetFPS)
	r.targetFPS = targetFPS
}

func (raylib) drawing(w *world) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		drawFood(w.food, w.cellWidth, w.cellWidth, rl.Red)
		drawSnake(w.snake, w.cellWidth, w.cellHeight, rl.DarkPurple, rl.Purple)
		drawScore(w.score)
		rl.EndDrawing()
	}
}

func (raylib) closeView() {
	rl.CloseWindow()
}

func drawSnake(s snake, width, height int32, c1, c2 rl.Color) {
	head := s.body[0]
	switch s.dir {
	case left:
		rl.DrawRectangleGradientH(head.x*width, head.y*height, width, height, c1, c2)
	case down:
		rl.DrawRectangleGradientV(head.x*width, head.y*height, width, height, c2, c1)
	case up:
		rl.DrawRectangleGradientV(head.x*width, head.y*height, width, height, c1, c2)
	case right:
		rl.DrawRectangleGradientH(head.x*width, head.y*height, width, height, c2, c1)
	}
	for _, seg := range s.body[1:] {
		rl.DrawRectangle(seg.x*width, seg.y*height, width, height, c2)
	}
}

func drawFood(f food, width, height int32, c rl.Color) {
	rl.DrawRectangle(f.x*width, f.y*height, width, height, c)
}

func drawScore(score int32) {
	headerText := fmt.Sprintf("SCORE: %d", score)
	rl.DrawText(headerText, 8, 8, rl.GetFontDefault().BaseSize, rl.Black)
}

func (r raylib) capturingInput(c chan<- control) {
	tick := time.NewTicker(time.Second / time.Duration(r.targetFPS))
	for range tick.C {
		if rl.IsKeyDown(rl.KeyUp) {
			c <- up
		} else if rl.IsKeyDown(rl.KeyLeft) {
			c <- left
		} else if rl.IsKeyDown(rl.KeyDown) {
			c <- down
		} else if rl.IsKeyDown(rl.KeyRight) {
			c <- right
		}
	}
	tick.Stop()
}

func (r raylib) handlingInput(c <-chan control, w *world) {
	for input := range c {
		w.snake.dir = input
	}
}
