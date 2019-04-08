package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := ViewData{
			Title:   "World Cup",
			Message: "Fifa will never regret it",
		}

		tmpl := template.Must(template.New("data").Parse(`
               
                     <div>
                    <h1>{{ .Title}}</h1>
                     <p>{{.Message}}</p>
                </div>
			`))
		tmpl.Execute(writer, data)
	})

	fmt.Println("Server listening")
	http.ListenAndServe(":8181", nil)
}
