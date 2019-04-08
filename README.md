"# golang-Exmaple" 
package main
import (
	"fmt"
	"net/http"
)
func main() {
	resp, err := http.Get("https://lms.tut.tj")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil{
			break
		}
	}
}



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
