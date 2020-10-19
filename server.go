package main

import (
	"time"
	"net/http"
	"encoding/json"
	"log"
)

type info struct {
	T string `json:"time"`
}
//This function will be executed when performing an HTTP GET request to /time
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

