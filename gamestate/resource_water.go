package gamestate

type Water struct {
	Amount      int
	Delta       int
	Total       int
	IsAutomated bool
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
	return w.Amount
}

func (w *Water) SetAmount(amount int) {
	w.Amount = amount
	w.Total += amount
}

func (w *Water) IncrementAmount() {
	w.Amount += w.Delta
	w.Total += w.Delta
}

func (w *Water) ChangeAmount(amount int) {
	w.Amount += amount

	if amount > 0 {
		w.Total += amount
	}
}

func (w *Water) SetAutomated(automated bool) {
	w.IsAutomated = automated
}

func (w *Water) GetIsAutomated() bool {
	return w.IsAutomated
}

func (w *Water) Tick(gameState *GameState) {
	if w.IsAutomated {
		w.Amount += w.Delta
		w.Total += w.Delta
	}

	if w.Amount > 0 {
		tool := gameState.GetTool("search-water")
		if tool == nil {
			gameState.Tools = append(gameState.Tools, &SearchWater{})
		}
	}
}
