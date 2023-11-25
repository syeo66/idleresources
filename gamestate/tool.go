package gamestate

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Effect struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	Property string `json:"property"`
	Amount   int    `json:"amount"`
}

type tool struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Costs     []Resource `json:"costs"`
	Effects   []Effect   `json:"effects"`
	IsEnabled bool       `json:"is_enabled"`
}

func NewTool(id string) *tool {
	path := "gamestate/tools/" + id + ".json"
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer jsonFile.Close()

	var newTool tool = tool{}
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &newTool)

	if err != nil {
		fmt.Println(err)
	}

	return &newTool
}

func (s *tool) GetId() string {
	return s.Id
}

func (s *tool) GetName() string {
	return s.Name
}

func (s *tool) GetCosts() []Resource {
	return s.Costs
}

func (s *tool) GetIsEnabled(gameState *GameState) bool {
	return s.IsEnabled
}

func (s *tool) Tick(gameState *GameState) {
}

func (s *tool) Act(gameState *GameState) {
	for _, effect := range s.Effects {
		switch effect.Type {
		case "resource":
			resource := gameState.GetResource(effect.Id)

			if resource == nil {
				resource = NewResource(effect.Id)
				gameState.Resources = append(gameState.Resources, resource)
			}

			switch effect.Property {
			case "amount":
				resource.ChangeAmount(effect.Amount)
			case "delta":
				resource.IncrementDelta(effect.Amount)
			}
		}
	}

	newCosts := []Resource{}
	for _, cost := range s.Costs {
		delta := cost.GetAmount()
		resource := gameState.GetResource(cost.GetId())
		resource.ChangeAmount(-delta)
		cost.ChangeAmount(resource.GetDelta() / 2)
		newCosts = append(newCosts, cost)
	}
	s.Costs = newCosts

	gameState.Compute()

	gameState.C <- *gameState
}

func (s *tool) Compute(gameState *GameState) {
	s.IsEnabled = true

	for _, cost := range s.Costs {
		available := gameState.GetResourceAmount(cost.GetId())
		if available < cost.GetAmount() {
			s.IsEnabled = false
			break
		}
	}
}
