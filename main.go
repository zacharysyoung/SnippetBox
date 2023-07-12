package main

import (
	"log"
	"net/http"
)

// home returns home.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from SnippetBox"))
}

// snippetView returns a specific snippet.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// snippetCreate creates a new snippet.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

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
