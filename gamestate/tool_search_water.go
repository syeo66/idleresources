package gamestate

type searchWater struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Costs     []Resource `json:"costs"`
	IsEnabled bool       `json:"is_enabled"`
}

func NewSearchWater() *searchWater {
	tool := &searchWater{
		Id:   "search-water",
		Name: "Search Water",
		Costs: []Resource{
			NewWater(),
		},
	}

	return tool
}

func (s *searchWater) GetId() string {
	return s.Id
}

func (s *searchWater) GetName() string {
	return s.Name
}

func (s *searchWater) GetCosts() []Resource {
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
	water.ChangeAmount(-1)
	gameState.Compute()

	gameState.C <- *gameState
}

func (s *searchWater) Compute(gameState *GameState) {
	s.IsEnabled = gameState.GetResourceAmount("water") > 0
}
