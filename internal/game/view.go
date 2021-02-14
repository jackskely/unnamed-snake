package game

type view interface {
	initView(width, height int32, title string, targetFPS int32)
	drawing(w *world)
	closeView()
}
