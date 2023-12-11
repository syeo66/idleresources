package gamestate

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ScienceEnables struct {
	Tools   []string `json:"tools"`
	Science []string `json:"science"`
}

type Science struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Enables  ScienceEnables `json:"enables"`
	Requires map[string]int `json:"requires"`
}

func NewScience(id string) *Science {
	path := "gamestate/science/" + id + ".json"
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer jsonFile.Close()

	var newScience Science = Science{}
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &newScience)

	if err != nil {
		fmt.Println(err)
	}

	return &newScience
}
