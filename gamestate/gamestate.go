package gamestate

type Resource interface {
	Id() string
	Name() string
	Amount() int
	Delta() int

	SetAmount(int)
	IncrementAmount(int)

	Tick(gameState *GameState)
}

type Tool interface {
	Id() string
	Name() string
	Costs() []Resource

	CanAfford(gameState *GameState) bool

	Tick(gameState *GameState)
}

type GameState struct {
	Tools     []Tool
	Resources []Resource
}

func (g *GameState) GetResource(Id string) Resource {
	for i, resource := range g.Resources {
		if resource.Id() == Id {
			return g.Resources[i]
		}
	}

	return nil
}

func (g *GameState) GetTool(Id string) Tool {
	for i, tool := range g.Tools {
		if tool.Id() == Id {
			return g.Tools[i]
		}
	}

	return nil
}

func (g *GameState) Tick() {
	for _, resource := range g.Resources {
		resource.Tick(g)
	}

	for _, tool := range g.Tools {
		tool.Tick(g)
	}
}
