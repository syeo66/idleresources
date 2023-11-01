package main

import (
	"html/template"
	"log"
	"net/http"
)

type Resource struct {
	Id     string
	Name   string
	Amount int
	Delta  int
}

type GameState struct {
	Resources []Resource
}

var gameState = GameState{
	Resources: []Resource{
		{Id: "water", Name: "Water", Amount: 0, Delta: 0},
	},
}

func (g *GameState) GetResource(Id string) *Resource {
	for i, resource := range g.Resources {
		if resource.Id == Id {
			return &g.Resources[i]
		}
	}

	return nil
}

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/resource.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, resource any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", resource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func water(w http.ResponseWriter, r *http.Request) {
	resource := gameState.GetResource("water")
	resource.Amount += 1
	renderTemplate(w, "resource", resource)
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", gameState)
}

func main() {
	fileServer := http.StripPrefix("/css", http.FileServer(http.Dir("./static/css")))
	http.Handle("/css/", fileServer)

	http.HandleFunc("/", index)
	http.HandleFunc("/water", water)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
