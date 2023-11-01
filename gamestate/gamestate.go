package gamestate

type Resource struct {
	Id     string
	Name   string
	Amount int
	Delta  int
}

type Tool struct {
	Id   string
	Name string
	Cost []Resource
}

type GameState struct {
	Tools     []Tool
	Resources []Resource
}

func (g *GameState) GetResource(Id string) *Resource {
	for i, resource := range g.Resources {
		if resource.Id == Id {
			return &g.Resources[i]
		}
	}

	return nil
}
