package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("say hello-wrold")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		fmt.Println("passou aqui")
	}
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
