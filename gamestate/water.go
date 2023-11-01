package gamestate

type Water struct {
	amount int
	delta  int
}

func (w *Water) Id() string {
	return "water"
}

func (w *Water) Name() string {
	return "Water"
}

func (w *Water) Amount() int {
	return w.amount
}

func (w *Water) Delta() int {
	return w.delta
}

func (w *Water) SetAmount(amount int) {
	w.amount = amount
}

func (w *Water) IncrementAmount(amount int) {
	w.amount += amount
}

func (w *Water) Tick(gameState *GameState) {
	w.amount += w.delta
}
