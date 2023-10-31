package main

import (
  "html/template"
  "net/http"
)

func main() {
  tmpl := template.Must(template.ParseFiles("templates/index.html"))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, nil)
  })

  http.ListenAndServe(":8080", nil)
}

