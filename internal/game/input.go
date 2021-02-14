package game

type inputHandler interface {
	capturingInput(c chan<- control)
	handlingInput(c <-chan control, w *world)
}
