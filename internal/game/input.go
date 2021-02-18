package game

type inputHandler interface {
	capturingInput(c chan<- direction)
	handlingInput(c <-chan direction, w *world)
}
