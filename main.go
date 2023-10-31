package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    err := tmpl.Execute(w, nil)

    if err != nil {
      fmt.Print(err)
    }
  }

func main() {
  http.HandleFunc("/", index)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

