package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := "Go Template"
		templ, _ := template.New("data").Parse("<h1 style=\"color: blue\">{{ .}}</h1>")
		templ.Execute(writer, data)
	})

	fmt.Println("Server listening...")
	http.ListenAndServe(":8181", nil)
}
