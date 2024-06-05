package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootGetHandler)
	mux.HandleFunc("POST /", rootPostHandler)

	mux.HandleFunc("GET /one", oneGetHandler)
	mux.HandleFunc("POST /one", onePostHandler)

	mux.HandleFunc("GET /two", twoGetHandler)
	mux.HandleFunc("POST /two", twoPostHandler)

	// regex check is impemented in the handler function
	mux.HandleFunc("GET /p/", pGetHandler)
	mux.HandleFunc("POST /p/", pPostHandler)

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
	pattern := regexp.MustCompile("^/p/([1-9][0-9]*)/c/([A-Z][A-Z])")
	if pattern.MatchString(r.URL.Path) {
		matches := pattern.FindStringSubmatch(r.URL.Path)
		id := matches[1]
		cc := matches[2]

		r.ParseForm()

		fmt.Fprintf(w, "get /p/%s/c/%s a='%s'", id, cc, r.Form.Get("a"))
	}
}

func pPostHandler(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile("^/p/([1-9][0-9]*)/c/([A-Z][A-Z])")
	if pattern.MatchString(r.URL.Path) {
		matches := pattern.FindStringSubmatch(r.URL.Path)
		id := matches[1]
		cc := matches[2]

		r.ParseForm()

		fmt.Fprintf(w, "post /p/%s/c/%s a='%s'", id, cc, r.Form.Get("a"))
	}
}
