package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func water(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/water.html"))
	err := tmpl.Execute(w, nil)

	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/water", water)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
