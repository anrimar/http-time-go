package main

import (
	"time"
	"net/http"
	"encoding/json"
	"log"
	"io"
)

type info struct {
	T string `json:"time"`
}

func timeHandler (w http.ResponseWriter, r *http.Request) {
	resp := info{time.Now().Format(time.RFC3339)}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
	log.Println(err)
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}

