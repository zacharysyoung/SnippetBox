package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	const hostPort = ":4000"
	log.Print("Starting server on", hostPort)
	err := http.ListenAndServe(hostPort, mux)
	log.Fatal(err)
}
