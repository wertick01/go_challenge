package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.Get("https://go.dev/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.Cookies())

	
	defer req.Body.Close()
}