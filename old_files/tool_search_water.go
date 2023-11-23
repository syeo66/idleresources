package gamestate

type searchWater struct {
	Id        string              `json:"id"`
	Name      string              `json:"name"`
	Costs     []ResourceInterface `json:"costs"`
	IsEnabled bool                `json:"is_enabled"`
}

func NewSearchWater() *searchWater {
	tool := &searchWater{
		Id:   "search-water",
		Name: "Search Water",
		Costs: []ResourceInterface{
			NewWater(),
		},
	}

	tool.Costs[0].SetAmount(1)

	return tool
}

func (s *searchWater) GetId() string {
	return s.Id
}

func (s *searchWater) GetName() string {
	return s.Name
}

func (s *searchWater) GetCosts() []ResourceInterface {
	return s.Costs
}

func (s *searchWater) GetIsEnabled(gameState *GameState) bool {
	return s.IsEnabled
}

func (s *searchWater) Tick(gameState *GameState) {
}

func (s *searchWater) Act(gameState *GameState) {
	water := gameState.GetResource("water")
	water.IncrementDelta(1)

	for _, cost := range s.Costs {
		delta := cost.GetAmount()
		resource := gameState.GetResource(cost.GetId())
		resource.ChangeAmount(-delta)
		cost.ChangeAmount(resource.GetDelta() / 2)
	}

	gameState.Compute()

	gameState.C <- *gameState
}

func (s *searchWater) Compute(gameState *GameState) {
	s.IsEnabled = true

	for _, cost := range s.Costs {
		available := gameState.GetResourceAmount(cost.GetId())
		if available < cost.GetAmount() {
			s.IsEnabled = false
			break
		}
	}
}
