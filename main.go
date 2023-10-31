package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
  tmpl := template.Must(template.ParseFiles("templates/index.html"))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    err := tmpl.Execute(w, nil)

    if err != nil {
      fmt.Print(err)
    }
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}

