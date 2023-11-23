package gamestate

type searchIronOre struct {
	Id        string              `json:"id"`
	Name      string              `json:"name"`
	Costs     []ResourceInterface `json:"costs"`
	IsEnabled bool                `json:"is_enabled"`
}

func NewSearchIronOre() *searchIronOre {
	tool := &searchIronOre{
		Id:   "search-iron_ore",
		Name: "Search Iron Ore",
		Costs: []ResourceInterface{
			NewWater(),
		},
	}

	tool.Costs[0].SetAmount(4)

	return tool
}

func (s *searchIronOre) GetId() string {
	return s.Id
}

func (s *searchIronOre) GetName() string {
	return s.Name
}

func (s *searchIronOre) GetCosts() []ResourceInterface {
	return s.Costs
}

func (s *searchIronOre) GetIsEnabled(gameState *GameState) bool {
	return s.IsEnabled
}

func (s *searchIronOre) Tick(gameState *GameState) {
}

func (s *searchIronOre) Act(gameState *GameState) {
	ironOre := gameState.GetResource("iron_ore")
	if ironOre == nil {
		ironOre = NewIronOre()
		gameState.Resources = append(gameState.Resources, ironOre)
	}
	ironOre.IncrementDelta(1)

	for _, cost := range s.Costs {
		delta := cost.GetAmount()
		resource := gameState.GetResource(cost.GetId())
		resource.ChangeAmount(-delta)
		cost.ChangeAmount(resource.GetDelta() / 2)
	}

	gameState.Compute()

	gameState.C <- *gameState
}

func (s *searchIronOre) Compute(gameState *GameState) {
	s.IsEnabled = true

	for _, cost := range s.Costs {
		available := gameState.GetResourceAmount(cost.GetId())
		if available < cost.GetAmount() {
			s.IsEnabled = false
			break
		}
	}
}
