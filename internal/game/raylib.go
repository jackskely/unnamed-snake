package game

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type raylib struct {
	title         string
	width, height int32
	targetFPS     int32
}

func newRaylib(width, height int32, title string, targetFPS int32) raylib {
	return raylib{
		title:     title,
		width:     width,
		height:    height,
		targetFPS: targetFPS,
	}
}

func (r raylib) initView() {
	rl.SetTraceLog(rl.LogNone)
	rl.InitWindow(r.width, r.height, r.title)
	rl.SetTargetFPS(r.targetFPS)
}

func (raylib) drawing(worldC <-chan world) {
	w := <-worldC
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		drawFood(w.food, w.cellWidth, w.cellHeight, rl.Red)
		drawSnake(w.snake, w.cellWidth, w.cellHeight, rl.DarkPurple, rl.Purple)
		drawScore(w.score)

		rl.EndDrawing()
		select {
		case snap := <-worldC:
			if snap.snake.isAlive {
				w = snap
			}
		default:
		}
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

func (r raylib) handleInput() <-chan direction {
	c := make(chan direction)

	go func() {
		defer close(c)
		tick := time.NewTicker(time.Second / time.Duration(r.targetFPS))
		for range tick.C {
			if rl.IsKeyPressed(rl.KeyUp) {
				c <- up
			}
			if rl.IsKeyPressed(rl.KeyLeft) {
				c <- left
			}
			if rl.IsKeyPressed(rl.KeyDown) {
				c <- down
			}
			if rl.IsKeyPressed(rl.KeyRight) {
				c <- right
			}
			if rl.IsKeyPressed(rl.KeyKp5) {
				break
			}
		}
		tick.Stop()
	}()

	return c
}
