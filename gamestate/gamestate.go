package gamestate

type Resource interface {
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
}

type Tool interface {
	Id() string
	Name() string
	Costs() []Resource

	IsEnabled(gameState *GameState) bool

	Tick(gameState *GameState)
}

type GameState struct {
	Tools     []Tool         `json:"tools"`
	Resources []Resource     `json:"resources"`
	C         chan GameState `json:"-"`
}

func (g *GameState) GetResource(Id string) Resource {
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
		if tool.Id() == Id {
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
		g.C <- *g
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

func (g *GameState) Init() {
	g.C = make(chan GameState)
	g.C <- *g
}
