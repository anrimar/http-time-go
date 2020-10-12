package main

import (
	"time"
	"net/http"
	"encoding/json"
)

type info struct {
	t string
}

func timeHandler (w http.ResponseWriter, r *http.Request) {
	//resp := info{t: time.Now().Format(time.RFC3339)}
	t := time.Now().Format(time.RFC3339)
	json.NewEncoder(w).Encode(t)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}
