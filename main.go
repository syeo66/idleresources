package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/syeo66/idleresources/gamestate"
)

var gameState = gamestate.GameState{
	Resources: []gamestate.Resource{
		{Id: "water", Name: "Water", Amount: 0, Delta: 0},
	},
}

var templatePaths = []string{
	"templates/index.html", "templates/resource.html", "templates/tools.html",
}
var templates = template.Must(template.ParseFiles(templatePaths...))

func renderTemplate(w http.ResponseWriter, tmpl string, resource any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", resource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	resourceName := r.URL.Path[len("/"):]

	resource := gameState.GetResource(resourceName)
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
	http.HandleFunc("/water", resourceHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
