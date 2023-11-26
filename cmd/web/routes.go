package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infolog.Printf("%s-%s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
func (app *application) routes() http.Handler {
	standardmiddleware := alice.New(app.logRequest)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.getuserinput))
	mux.Post("/showcomic", http.HandlerFunc(app.indexHandler)) //pat does not allow us to use handlers directly
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardmiddleware.Then(mux)
}
