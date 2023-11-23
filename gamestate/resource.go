package gamestate

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Resource struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Amount      int              `json:"amount"`
	Delta       int              `json:"delta"`
	Total       int              `json:"total"`
	IsAutomated bool             `json:"is_automated"`
	EnableTools map[int][]string `json:"enable_tools"`
}

func NewResource(id string) *Resource {
	path := "gamestate/resources/" + id + ".json"
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer jsonFile.Close()

	var newResource Resource = Resource{}
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &newResource)

	if err != nil {
		fmt.Println(err)
	}

	return &newResource
}

func (w *Resource) GetId() string {
	return w.Id
}

func (w *Resource) GetName() string {
	return w.Name
}

func (w *Resource) GetDelta() int {
	return w.Delta
}

func (w *Resource) SetDelta(delta int) {
	w.Delta = delta
}

func (w *Resource) IncrementDelta(delta int) {
	w.Delta += delta
}

func (w *Resource) GetAmount() int {
	return w.Amount
}

func (w *Resource) SetAmount(amount int) {
	w.Amount = amount
	w.Total += amount
}

func (w *Resource) IncrementAmount() {
	w.Amount += w.Delta
	w.Total += w.Delta
}

func (w *Resource) ChangeAmount(amount int) {
	w.Amount += amount

	if amount > 0 {
		w.Total += amount
	}
}

func (w *Resource) SetAutomated(automated bool) {
	w.IsAutomated = automated
}

func (w *Resource) GetIsAutomated() bool {
	return w.IsAutomated
}

func (w *Resource) Tick(gameState *GameState) {
	if w.IsAutomated {
		w.Amount += w.Delta
		w.Total += w.Delta
	}

	gameState.Compute()
}

func (w *Resource) Compute(gameState *GameState) {
	for amount, tools := range w.EnableTools {
		if w.Amount > amount {
			for _, toolId := range tools {
				tool := gameState.GetTool(toolId)
				if tool == nil {
					newTool := NewTool(toolId)
					gameState.Tools = append(gameState.Tools, newTool)
				}
			}
		}
	}
}
