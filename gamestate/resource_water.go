package gamestate

type water struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Delta       int    `json:"delta"`
	Total       int    `json:"total"`
	IsAutomated bool   `json:"is_automated"`
}

func NewWater() Resource {
	return &water{
		Id:          "water",
		Name:        "Water",
		Amount:      0,
		Delta:       1,
		Total:       0,
		IsAutomated: false,
	}
}

func (w *water) GetId() string {
	return w.Id
}

func (w *water) GetName() string {
	return w.Name
}

func (w *water) GetDelta() int {
	return w.Delta
}

func (w *water) SetDelta(delta int) {
	w.Delta = delta
}

func (w *water) IncrementDelta(delta int) {
	w.Delta += delta
}

func (w *water) GetAmount() int {
	return w.Amount
}

func (w *water) SetAmount(amount int) {
	w.Amount = amount
	w.Total += amount
}

func (w *water) IncrementAmount() {
	w.Amount += w.Delta
	w.Total += w.Delta
}

func (w *water) ChangeAmount(amount int) {
	w.Amount += amount

	if amount > 0 {
		w.Total += amount
	}
}

func (w *water) SetAutomated(automated bool) {
	w.IsAutomated = automated
}

func (w *water) GetIsAutomated() bool {
	return w.IsAutomated
}

func (w *water) Tick(gameState *GameState) {
	if w.IsAutomated {
		w.Amount += w.Delta
		w.Total += w.Delta
	}

	gameState.Compute()
}

func (w *water) Compute(gameState *GameState) {
	if w.Amount > 0 {
		tool := gameState.GetTool("search-water")
		if tool == nil {
			searchWater := NewSearchWater()
			gameState.Tools = append(gameState.Tools, searchWater)
		}
	}

	if w.Amount > 3 {
		tool := gameState.GetTool("search-coal")
		if tool == nil {
			searchCoal := NewSearchCoal()
			gameState.Tools = append(gameState.Tools, searchCoal)
		}
	}

	if w.Amount > 6 {
		tool := gameState.GetTool("search-iron_ore")
		if tool == nil {
			searchIronOre := NewSearchIronOre()
			gameState.Tools = append(gameState.Tools, searchIronOre)
		}
	}
}
