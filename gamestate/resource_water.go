package gamestate

type Water struct {
	amount      int
	Delta       int
	total       int
	isAutomated bool
}

func (w *Water) Id() string {
	return "water"
}

func (w *Water) Name() string {
	return "Water"
}

func (w *Water) GetDelta() int {
	return w.Delta
}

func (w *Water) SetDelta(delta int) {
	w.Delta = delta
}

func (w *Water) IncrementDelta(delta int) {
	w.Delta += delta
}

func (w *Water) GetAmount() int {
	return w.amount
}

func (w *Water) SetAmount(amount int) {
	w.amount = amount
	w.total += amount
}

func (w *Water) IncrementAmount() {
	w.amount += w.Delta
	w.total += w.Delta
}

func (w *Water) ChangeAmount(amount int) {
	w.amount += amount

	if amount > 0 {
		w.total += amount
	}
}

func (w *Water) SetAutomated(automated bool) {
	w.isAutomated = automated
}

func (w *Water) IsAutomated() bool {
	return w.isAutomated
}

func (w *Water) Tick(gameState *GameState) {
	if w.isAutomated {
		w.amount += w.Delta
		w.total += w.Delta
	}

	if w.amount > 0 {
		tool := gameState.GetTool("search-water")
		if tool == nil {
			gameState.Tools = append(gameState.Tools, &SearchWater{})
		}
	}
}
