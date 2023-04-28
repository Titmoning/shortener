package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/unshorten", unshortenHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Fatalf("Error wrong method")
	}
	var u shortenRequst
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := shortenResponse{
		URL:      u.URL,
		ShortURL: "text2",
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)

	fmt.Printf("Got req %v \n", u)
}

type shortenRequst struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
}

func unshortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Fatalf("Error wrong method")
	}
	shortURL := r.URL.Query().Get("short_url")
	if shortURL != "" {
	}
	resp := unshortenResponse{
		URL:      "text",
		ShortURL: shortURL,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	_ = err
	fmt.Printf("Got req %v \n", shortURL)

}

type unshortenResponse struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
}
