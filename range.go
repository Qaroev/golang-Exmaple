package main

import (
	"net/http"
	"html/template"
	"fmt"
)

type ViewDatas struct {
	Title string
	Users []string
}

func main() {
	data := ViewDatas{
		Title: "Users List",
		Users: []string{"Tom", "Bob", "Sam"},
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		templ, _ := template.ParseFiles("range.html")
		templ.Execute(writer, data)
	})

	fmt.Println("Server listening...")
	http.ListenAndServe(":8282", nil)
}
