package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootGetHandler)
	mux.HandleFunc("POST /", rootPostHandler)

	mux.HandleFunc("GET /one", oneGetHandler)
	mux.HandleFunc("POST /one", onePostHandler)

	mux.HandleFunc("GET /two", twoGetHandler)
	mux.HandleFunc("POST /two", twoPostHandler)

	// TODO: implement regex
	mux.HandleFunc("GET /p/{id}/c/{cc}", pGetHandler)
	mux.HandleFunc("POST /p/{id}/c/{cc}", pPostHandler)

	http.ListenAndServe(":4567", mux)
}

func rootGetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "get / a='%s'", r.Form.Get("a"))
}

func rootPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "post / a='%s'", r.Form.Get("a"))
}

func oneGetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "get /one a='%s'", r.Form.Get("a"))
}

func onePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "post /one a='%s'", r.Form.Get("a"))
}

func twoGetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "get /two a='%s' b='%s'", r.Form.Get("a"), r.Form.Get("b"))
}

func twoPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "post /two a='%s' b='%s'", r.Form.Get("a"), r.Form.Get("b"))
}

func pGetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.PathValue("id")
	cc := r.PathValue("cc")
	fmt.Fprintf(w, "get /p/%s/c/%s a='%s'", id, cc, r.Form.Get("a"))
}

func pPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.PathValue("id")
	cc := r.PathValue("cc")
	fmt.Fprintf(w, "get /p/%s/c/%s a='%s'", id, cc, r.Form.Get("a"))
}
