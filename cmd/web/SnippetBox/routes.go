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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(get, "/", dynamic.ThenFunc(app.home))
	router.Handler(get, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(get, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	router.Handler(post, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
