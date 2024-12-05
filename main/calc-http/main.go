package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mdw-cohort-c/calc-apps/handlers"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	err := http.ListenAndServe("localhost:8080", handlers.NewRouter(logger))
	if err != nil {
		log.Fatal(err)
	}
}
