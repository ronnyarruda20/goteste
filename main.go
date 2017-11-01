package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	rotas := mux.NewRouter().StrictSlash(true)
	rotas.HandleFunc("/", getMembros).Methods("GET")

	log.Fatal(http.ListenAndServe(port, rotas))

}

func getMembros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// u := User{Id: "US123", Balance: 8}

	json.NewEncoder(w).Encode("u")
}
