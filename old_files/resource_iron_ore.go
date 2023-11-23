package gamestate

type ironOre struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Delta       int    `json:"delta"`
	Total       int    `json:"total"`
	IsAutomated bool   `json:"is_automated"`
}

func NewIronOre() ResourceInterface {
	return &ironOre{
		Id:          "iron_ore",
		Name:        "Iron Ore",
		Amount:      0,
		Delta:       0,
		Total:       0,
		IsAutomated: false,
	}
}

func (w *ironOre) GetId() string {
	return w.Id
}

func (w *ironOre) GetName() string {
	return w.Name
}

func (w *ironOre) GetDelta() int {
	return w.Delta
}

func (w *ironOre) SetDelta(delta int) {
	w.Delta = delta
}

func (w *ironOre) IncrementDelta(delta int) {
	w.Delta += delta
}

func (w *ironOre) GetAmount() int {
	return w.Amount
}

func (w *ironOre) SetAmount(amount int) {
	w.Amount = amount
	w.Total += amount
}

func (w *ironOre) IncrementAmount() {
	w.Amount += w.Delta
	w.Total += w.Delta
}

func (w *ironOre) ChangeAmount(amount int) {
	w.Amount += amount

	if amount > 0 {
		w.Total += amount
	}
}

func (w *ironOre) SetAutomated(automated bool) {
	w.IsAutomated = automated
}

func (w *ironOre) GetIsAutomated() bool {
	return w.IsAutomated
}

func (w *ironOre) Tick(gameState *GameState) {
	if w.IsAutomated {
		w.Amount += w.Delta
		w.Total += w.Delta
	}

	gameState.Compute()
}

func (w *ironOre) Compute(gameState *GameState) {
}
