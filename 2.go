package main
import (
	"fmt"
	"net/http"
	"io"
	"os"
)
func main() {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://somon.tj", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")   // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0")   // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}