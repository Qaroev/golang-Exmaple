package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/product/{id:[0-9]+}", productHandlerss)

	http.Handle("/", router)

	fmt.Println("Server listening...")
	http.ListenAndServe(":8080",nil)
}

func productHandlerss(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response := fmt.Sprintf("Product %s", id)

	fmt.Fprintf(w, response)
}
