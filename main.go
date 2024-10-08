package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", pingHandler)

	log.Println("Server Started and listening on :8181")
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping Handler hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}
