package game

type inputHandler interface {
	handleInput() <-chan direction
}
