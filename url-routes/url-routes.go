package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootGetHandler)

	mux.HandleFunc("GET /people", peopleGetHandler)
	mux.HandleFunc("POST /people", peoplePostHandler)

	mux.HandleFunc("GET /people/", peopleIdGetHandler)
	mux.HandleFunc("POST /people/{id}", peopleIdPostHandler)
	// using regex to match the url
	mux.HandleFunc("GET /people/{id}/country/{cc}", personCountryHandler)
	mux.HandleFunc("GET /country/", countryHandler)

	http.ListenAndServe(":4567", mux)
}

func rootGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "got /")
		return
	}
	http.NotFound(w, r)
}

func peopleGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/people" {
		fmt.Fprint(w, "got /people")
		return
	}
	http.NotFound(w, r)
}

func peoplePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/people" {
		fmt.Fprint(w, "posted to /people")
		return
	}
	http.NotFound(w, r)
}

func peopleIdGetHandler(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`^/people/(\d+)$`)
	if !pattern.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}
	matches := pattern.FindStringSubmatch(r.URL.Path)
	id := matches[1]

	fmt.Fprintf(w, "got person %s", id)
}

func peopleIdPostHandler(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`^/people/(\d+)$`)
	if !pattern.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}
	matches := pattern.FindStringSubmatch(r.URL.Path)
	id := matches[1]

	fmt.Fprintf(w, "posted person %s", id)
}

func countryHandler(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`^/country/([A-Z]{2})$`)
	if !pattern.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}
	matches := pattern.FindStringSubmatch(r.URL.Path)
	cc := matches[1]

	fmt.Fprintf(w, "got country %s", cc)
}

func personCountryHandler(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile(`^/people/([1-9][0-9]*)/country/([A-Z]{2})$`)
	if !pattern.MatchString(r.URL.Path) {
		http.NotFound(w, r)
		return
	}
	matches := pattern.FindStringSubmatch(r.URL.Path)
	id := matches[1]
	cc := matches[2]

	fmt.Fprintf(w, "got person %s, country %s", id, cc)
}
