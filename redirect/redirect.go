package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", redirectHandler)
	mux.HandleFunc("GET /here", hereHandler)

	http.ListenAndServe(":4567", mux)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// just /here works fine
	fullURL := "http://" + r.Host + "/here"
	http.Redirect(w, r, fullURL, http.StatusFound)
}

func hereHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "arrived")
}
