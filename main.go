package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/jbabineau/pomodoro-timer/templates"
)

func main() {

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/", indexHandler)

	log.Println("Server Started and listening on :8181")
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping Handler hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/hello/")
	if name == "" {
		name = "world"
	}

	page := templates.Hello(name)

	err := page.Render(r.Context(), w)

	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := templates.Index()

	err := page.Render(r.Context(), w)	

	if err != nil {	
		log.Fatal(err)
	}
}
