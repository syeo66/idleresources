package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/syeo66/idleresources/gamestate"
)

var gameState = gamestate.GameState{
	Resources: []gamestate.Resource{
		&gamestate.Water{},
	},
}

var templatePaths = []string{
	"templates/index.html",
	"templates/resource.html",
	"templates/resources.html",
	"templates/source.html",
	"templates/tools.html",
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
	resource.IncrementAmount(1)

	renderTemplate(w, "resources", gameState)
}

func resourcesView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "resources", gameState)
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", gameState)
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				gameState.Tick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	fileServer := http.StripPrefix("/css", http.FileServer(http.Dir("./static/css")))
	http.Handle("/css/", fileServer)

	http.HandleFunc("/", index)

	http.HandleFunc("/water", resourceHandler)

	http.HandleFunc("/resources", resourcesView)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
