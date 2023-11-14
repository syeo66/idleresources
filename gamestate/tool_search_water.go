package gamestate

type searchWater struct {
	Id    string     `json:"id"`
	Name  string     `json:"name"`
	Costs []Resource `json:"costs"`
}

func NewSearchWater() *searchWater {
	return &searchWater{
		Id:   "search-water",
		Name: "Search Water",
		Costs: []Resource{
			NewWater(),
		},
	}
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

func (s *searchWater) IsEnabled(gameState *GameState) bool {
	return gameState.GetResourceAmount("water") > 0
}

func (s *searchWater) Tick(gameState *GameState) {
}

func (s *searchWater) Act(gameState *GameState) {
}
