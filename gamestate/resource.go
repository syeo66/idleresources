package gamestate

import (
	"encoding/json"
	"io"
	"os"
)

type resource struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Delta       int    `json:"delta"`
	Total       int    `json:"total"`
	IsAutomated bool   `json:"is_automated"`
}

func NewResource(path string) (*resource, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	var newResource resource
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &newResource)

	if err != nil {
		return nil, err
	}

	return &newResource, nil
}

func (w *resource) GetId() string {
	return w.Id
}

func (w *resource) GetName() string {
	return w.Name
}

func (w *resource) GetDelta() int {
	return w.Delta
}

func (w *resource) SetDelta(delta int) {
	w.Delta = delta
}

func (w *resource) IncrementDelta(delta int) {
	w.Delta += delta
}

func (w *resource) GetAmount() int {
	return w.Amount
}

func (w *resource) SetAmount(amount int) {
	w.Amount = amount
	w.Total += amount
}

func (w *resource) IncrementAmount() {
	w.Amount += w.Delta
	w.Total += w.Delta
}

func (w *resource) ChangeAmount(amount int) {
	w.Amount += amount

	if amount > 0 {
		w.Total += amount
	}
}

func (w *resource) SetAutomated(automated bool) {
	w.IsAutomated = automated
}

func (w *resource) GetIsAutomated() bool {
	return w.IsAutomated
}

func (w *resource) Tick(gameState *GameState) {
	if w.IsAutomated {
		w.Amount += w.Delta
		w.Total += w.Delta
	}

	gameState.Compute()
}

func (w *resource) Compute(gameState *GameState) {
}
