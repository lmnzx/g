package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /show/all", infoHeaderHandler)
	mux.HandleFunc("GET /lost", lostHandler)
	mux.HandleFunc("GET /teapot", teaPotHandler)
	mux.HandleFunc("GET /setter", setHeaderHandler)

	http.ListenAndServe(":4567", mux)
}

func infoHeaderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PATH_INFO = %s\n", r.URL.Path)
	fmt.Fprintf(w, "QUERY_STRING = %s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "REMOTE_ADDR = %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "REQUEST_METHOD = %s\n", r.Method)
	fmt.Fprintf(w, "REQUEST_URI = %s\n", r.RequestURI)
	fmt.Fprintf(w, "SERVER_NAME = %s\n", r.Host)
	fmt.Fprintf(w, "SERVER_PORT = %s\n", r.URL.Port())
	fmt.Fprintf(w, "HTTP_HOST = %s\n", r.Host)
	fmt.Fprintf(w, "HTTP_USER_AGENT = %s\n", r.UserAgent())
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
