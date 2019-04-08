package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "index.html")
	})
	http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "about.html")
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Index Page")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8081", nil)
}
