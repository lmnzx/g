package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /give", setCookieHandler)
	mux.HandleFunc("GET /delete", deleteCookieHandler)
	mux.HandleFunc("GET /show", getCookieHandler)

	http.ListenAndServe(":4567", mux)
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "ok",
		Value:    "yeah",
		Path:     "/",
		Expires:  time.Now().Add((time.Second * 60 * 60 * 24 * 3)),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("ok")

	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	}

	fmt.Fprint(w, cookie)
}

func deleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "ok",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-(time.Second * 60 * 60 * 24)),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
}
