package gamestate

type searchCoal struct {
	Id        string              `json:"id"`
	Name      string              `json:"name"`
	Costs     []ResourceInterface `json:"costs"`
	IsEnabled bool                `json:"is_enabled"`
}

func NewSearchCoal() *searchCoal {
	tool := &searchCoal{
		Id:   "search-coal",
		Name: "Search Coal",
		Costs: []ResourceInterface{
			NewWater(),
		},
	}

	tool.Costs[0].SetAmount(3)

	return tool
}

func (s *searchCoal) GetId() string {
	return s.Id
}

func (s *searchCoal) GetName() string {
	return s.Name
}

func (s *searchCoal) GetCosts() []ResourceInterface {
	return s.Costs
}

func (s *searchCoal) GetIsEnabled(gameState *GameState) bool {
	return s.IsEnabled
}

func (s *searchCoal) Tick(gameState *GameState) {
}

func (s *searchCoal) Act(gameState *GameState) {
	coal := gameState.GetResource("coal")
	if coal == nil {
		coal = NewCoal()
		gameState.Resources = append(gameState.Resources, coal)
	}
	coal.IncrementDelta(1)

	for _, cost := range s.Costs {
		delta := cost.GetAmount()
		resource := gameState.GetResource(cost.GetId())
		resource.ChangeAmount(-delta)
		cost.ChangeAmount(resource.GetDelta() / 2)
	}

	gameState.Compute()

	gameState.C <- *gameState
}

func (s *searchCoal) Compute(gameState *GameState) {
	s.IsEnabled = true

	for _, cost := range s.Costs {
		available := gameState.GetResourceAmount(cost.GetId())
		if available < cost.GetAmount() {
			s.IsEnabled = false
			break
		}
	}
}
