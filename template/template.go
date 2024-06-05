package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello/{name}", templateHandler)

	http.ListenAndServe(":4567", mux)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello %s</h1>", name)
}
