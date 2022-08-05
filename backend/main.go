package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	expectedAuthHeader := os.Getenv("AUTH_TOKEN")
	releasesDatabaseServiceUrl := os.Getenv("RELEASES_URL")

	if expectedAuthHeader == "" {
		panic("AUTH_TOKEN needs to be set!")
	} else {
		log.Printf("You've set the AUTH_TOKEN to: %s. But don't tell it to anyone!\n ", expectedAuthHeader)
	}

	if releasesDatabaseServiceUrl == "" {
		panic("RELEASES_URL needs to be set!")
	} else {
		log.Printf("You've set the RELEASES_URL to: %s. Should be there?\n ", releasesDatabaseServiceUrl)
	}

	http.HandleFunc("/releases", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader != expectedAuthHeader {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}

		if r.Method == "GET" {
			getForward(&w, releasesDatabaseServiceUrl+"/releases")
		}

		if r.Method == "POST" {
			postForward(&w, r, releasesDatabaseServiceUrl+"/releases")
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getForward(w *http.ResponseWriter, releaseServiceUrl string) {
	response, err := http.Get(releaseServiceUrl)

	if err != nil {
		writer := *w
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(*w, "Something went wrong "+err.Error())
		return
	}

	defer response.Body.Close()

	_, err = io.Copy(*w, response.Body)

	if err != nil {
		writer := *w
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(*w, "Something went wrong "+err.Error())
		return
	}
}

func postForward(w *http.ResponseWriter, r *http.Request, releaseServiceUrl string) {
	response, err := http.Post(releaseServiceUrl, "application/json", r.Body)

	if err != nil {
		writer := *w
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(*w, "Something went wrong "+err.Error())
		return
	}

	defer response.Body.Close()

	_, err = io.Copy(*w, response.Body)

	if err != nil {
		writer := *w
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(*w, "Something went wrong "+err.Error())
		return
	}
}
