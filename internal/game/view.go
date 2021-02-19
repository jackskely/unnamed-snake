package game

type view interface {
	initView()
	drawing(<-chan world)
	closeView()
}
