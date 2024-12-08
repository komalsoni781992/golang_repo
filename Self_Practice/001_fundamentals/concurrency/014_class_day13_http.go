package concurrency

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func TestHttp() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
