package gamestate

type SearchWater struct {
}

func (s *SearchWater) Id() string {
	return "search-water"
}

func (s *SearchWater) Name() string {
	return "Search Water"
}

func (s *SearchWater) Costs() []Resource {
	return []Resource{
		&Water{amount: 1},
	}
}

func (s *SearchWater) IsEnabled(gameState *GameState) bool {
	return gameState.GetResourceAmount("water") > 0
}

func (s *SearchWater) Tick(gameState *GameState) {
}

func (s *SearchWater) Act(gameState *GameState) {
}
