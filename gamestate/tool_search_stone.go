package gamestate

type searchStone struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Costs     []Resource `json:"costs"`
	IsEnabled bool       `json:"is_enabled"`
}

func NewSearchStone() *searchStone {
	tool := &searchStone{
		Id:   "search-stone",
		Name: "Search Stone",
		Costs: []Resource{
			NewWater(),
		},
	}

	tool.Costs[0].SetAmount(2)

	return tool
}

func (s *searchStone) GetId() string {
	return s.Id
}

func (s *searchStone) GetName() string {
	return s.Name
}

func (s *searchStone) GetCosts() []Resource {
	return s.Costs
}

func (s *searchStone) GetIsEnabled(gameState *GameState) bool {
	return s.IsEnabled
}

func (s *searchStone) Tick(gameState *GameState) {
}

func (s *searchStone) Act(gameState *GameState) {
	stone := gameState.GetResource("stone")
	if stone == nil {
		stone = NewStone()
		gameState.Resources = append(gameState.Resources, stone)
	}
	stone.IncrementDelta(1)

	for _, cost := range s.Costs {
		delta := cost.GetAmount()
		resource := gameState.GetResource(cost.GetId())
		resource.ChangeAmount(-delta)
		cost.ChangeAmount(resource.GetDelta() / 2)
	}

	gameState.Compute()

	gameState.C <- *gameState
}

func (s *searchStone) Compute(gameState *GameState) {
	s.IsEnabled = true

	for _, cost := range s.Costs {
		available := gameState.GetResourceAmount(cost.GetId())
		if available < cost.GetAmount() {
			s.IsEnabled = false
			break
		}
	}
}
