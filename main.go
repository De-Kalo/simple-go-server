package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")

	port := os.Getenv("PORT")
	fmt.Println("Running server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Status int
		Msg    string
	}
	w.Header().Set("Content-Type", "application/json")
	res := Result{0, "Still alive!"}
	fmt.Println(res)
	json.NewEncoder(w).Encode(res)
}
