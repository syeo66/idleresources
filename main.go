package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/syeo66/idleresources/gamestate"
)

var gameState = gamestate.GameState{
	Resources: []gamestate.Resource{
		&gamestate.Water{Delta: 1},
	},
	Tools: []gamestate.Tool{},
}

var templatePaths = []string{
	"templates/index.html",
	"templates/resource.html",
	"templates/resources.html",
	"templates/source.html",
	"templates/tool.html",
	"templates/tools.html",
}
var templates = template.Must(template.ParseFiles(templatePaths...))

func renderTemplate(w http.ResponseWriter, tmpl string, resource any) {
	err := templates.ExecuteTemplate(w, tmpl+".html", resource)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	resourceName := r.URL.Path[len("/search-"):]
	tool := gameState.GetTool("search-" + resourceName)

	if tool != nil {
		resource := gameState.GetResource(resourceName)
		resource.IncrementDelta(1)
		fmt.Printf("cost: %v", tool)

		for _, cost := range tool.Costs() {
			fmt.Println(cost.Id())
			gameState.GetResource(cost.Id()).ChangeAmount(-cost.GetAmount())
		}
	}

	renderTemplate(w, "tools", gameState)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	resourceName := r.URL.Path[len("/"):]

	resource := gameState.GetResource(resourceName)
	resource.IncrementAmount()

	renderTemplate(w, "resources", gameState)
}

func resourcesView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "resources", gameState)
}

func toolsView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "tools", gameState)
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

	http.HandleFunc("/search-water", searchHandler)

	http.HandleFunc("/resources", resourcesView)
	http.HandleFunc("/tools", toolsView)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
