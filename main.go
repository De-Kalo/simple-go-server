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
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Status int
		Msg    string
	}
	res := Result{0, "Still alive!"}
	fmt.Println(res)
	json.NewEncoder(w).Encode(res)
}
