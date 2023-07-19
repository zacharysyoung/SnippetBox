package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	const (
		get  = http.MethodGet
		post = http.MethodPost
	)

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(get, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(get, "/", app.home)
	router.HandlerFunc(get, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(get, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(post, "/snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
