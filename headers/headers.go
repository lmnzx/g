package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /show/all", infoHeaderHandler)
	mux.HandleFunc("GET /lost", lostHandler)
	mux.HandleFunc("GET /teapot", teaPotHandler)
	mux.HandleFunc("GET /setter", setHeaderHandler)

	http.ListenAndServe(":4567", mux)
}

// TODO: more info
func infoHeaderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request at %v\n", time.Now())
	for k, v := range r.Header {
		fmt.Fprintf(w, "%v: %v\n", k, v)
	}
}

func lostHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
}

func teaPotHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusTeapot)
}

func setHeaderHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Custom", "My-Own-Header")
	fmt.Fprint(w, "See Custom")
}
