package gamestate

type coal struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Delta       int    `json:"delta"`
	Total       int    `json:"total"`
	IsAutomated bool   `json:"is_automated"`
}

func NewCoal() ResourceInterface {
	return &coal{
		Id:          "coal",
		Name:        "Coal",
		Amount:      0,
		Delta:       0,
		Total:       0,
		IsAutomated: false,
	}
}

func (w *coal) GetId() string {
	return w.Id
}

func (w *coal) GetName() string {
	return w.Name
}

func (w *coal) GetDelta() int {
	return w.Delta
}

func (w *coal) SetDelta(delta int) {
	w.Delta = delta
}

func (w *coal) IncrementDelta(delta int) {
	w.Delta += delta
}

func (w *coal) GetAmount() int {
	return w.Amount
}

func (w *coal) SetAmount(amount int) {
	w.Amount = amount
	w.Total += amount
}

func (w *coal) IncrementAmount() {
	w.Amount += w.Delta
	w.Total += w.Delta
}

func (w *coal) ChangeAmount(amount int) {
	w.Amount += amount

	if amount > 0 {
		w.Total += amount
	}
}

func (w *coal) SetAutomated(automated bool) {
	w.IsAutomated = automated
}

func (w *coal) GetIsAutomated() bool {
	return w.IsAutomated
}

func (w *coal) Tick(gameState *GameState) {
	if w.IsAutomated {
		w.Amount += w.Delta
		w.Total += w.Delta
	}

	gameState.Compute()
}

func (w *coal) Compute(gameState *GameState) {
}
