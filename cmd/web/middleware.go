package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit and page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all PORT requests
func NoSurf(next http.Handler) http.Handler {
	cfrsHandler := nosurf.New(next)
	cfrsHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return cfrsHandler
}

// LoadSession loads and saves the session on every requests
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
