package main

import (
	"log"
	"net/http"

	"github.com/valp0/simple-api/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/hello", handlers.Hello)
	http.HandleFunc("/btc2usd", handlers.Bitcoin)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
