package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/releases", func(w http.ResponseWriter, r *http.Request) {
		// auth

		// GET

		// POST
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
