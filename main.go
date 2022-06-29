package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "9000"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("defaulting to port: %s", port)
	}

	a := App{}
	a.Init()

	http.Handle("/", a.Router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
