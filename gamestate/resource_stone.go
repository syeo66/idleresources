package gamestate

type stone struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Delta       int    `json:"delta"`
	Total       int    `json:"total"`
	IsAutomated bool   `json:"is_automated"`
}

func NewStone() Resource {
	return &stone{
		Id:          "stone",
		Name:        "Stone",
		Amount:      0,
		Delta:       0,
		Total:       0,
		IsAutomated: false,
	}
}

func (w *stone) GetId() string {
	return w.Id
}

func (w *stone) GetName() string {
	return w.Name
}

func (w *stone) GetDelta() int {
	return w.Delta
}

func (w *stone) SetDelta(delta int) {
	w.Delta = delta
}

func (w *stone) IncrementDelta(delta int) {
	w.Delta += delta
}

func (w *stone) GetAmount() int {
	return w.Amount
}

func (w *stone) SetAmount(amount int) {
	w.Amount = amount
	w.Total += amount
}

func (w *stone) IncrementAmount() {
	w.Amount += w.Delta
	w.Total += w.Delta
}

func (w *stone) ChangeAmount(amount int) {
	w.Amount += amount

	if amount > 0 {
		w.Total += amount
	}
}

func (w *stone) SetAutomated(automated bool) {
	w.IsAutomated = automated
}

func (w *stone) GetIsAutomated() bool {
	return w.IsAutomated
}

func (w *stone) Tick(gameState *GameState) {
	if w.IsAutomated {
		w.Amount += w.Delta
		w.Total += w.Delta
	}

	gameState.Compute()
}

func (w *stone) Compute(gameState *GameState) {
}
