package gamestate

type ResourceInterface interface {
	GetId() string
	GetName() string

	GetAmount() int
	SetAmount(int)
	IncrementAmount()
	ChangeAmount(int)

	GetDelta() int
	SetDelta(int)
	IncrementDelta(int)

	Tick(gameState *GameState)
	Compute(gameState *GameState)
}

type Tool interface {
	GetId() string
	GetName() string
	GetCosts() []Resource

	GetIsEnabled(gameState *GameState) bool

	Tick(gameState *GameState)
	Act(gameState *GameState)
	Compute(gameState *GameState)
}

type Technology interface {
	GetId() string
	GetName() string
	GetCosts() []Resource
	GetEnablesTechnology() []Technology

	GetIsEnabled(gameState *GameState) bool
}

type GameState struct {
	Tools        []Tool              `json:"tools"`
	Resources    []ResourceInterface `json:"resources"`
	Technologies []Technology        `json:"technologies"`
	C            chan GameState      `json:"-"`
}

func (g *GameState) GetResource(Id string) ResourceInterface {
	for i, resource := range g.Resources {
		if resource.GetId() == Id {
			return g.Resources[i]
		}
	}

	return nil
}

func (g *GameState) GetResourceAmount(Id string) int {
	resource := g.GetResource(Id)
	if resource == nil {
		return 0
	}

	return resource.GetAmount()
}

func (g *GameState) GetTool(Id string) Tool {
	for i, tool := range g.Tools {
		if tool.GetId() == Id {
			return g.Tools[i]
		}
	}

	return nil
}

func (g *GameState) HandleCommand(cmd map[string]interface{}) {
	cmdType := cmd["type"].(string)

	switch cmdType {
	case "collect":
		resource := g.GetResource(cmd["payload"].(string))
		if resource == nil {
			return
		}
		resource.IncrementAmount()
		g.Compute()
		g.C <- *g

	default:
		tool := g.GetTool(cmdType)
		if tool == nil || !tool.GetIsEnabled(g) {
			return
		}
		tool.Act(g)
	}
}

func (g *GameState) Tick() {
	for _, resource := range g.Resources {
		resource.Tick(g)
	}

	for _, tool := range g.Tools {
		tool.Tick(g)
	}
}

func (g *GameState) Compute() {
	for _, resource := range g.Resources {
		resource.Compute(g)
	}

	for _, tool := range g.Tools {
		tool.Compute(g)
	}
}

func (g *GameState) Init() {
	g.C = make(chan GameState)
	g.C <- *g
}
