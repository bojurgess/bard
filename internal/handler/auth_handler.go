package handler

import (
	"github.com/bojurgess/bard/internal/service"
	"net/http"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	redirect, state := service.SpotifyService.GenerateAuthUrl()
	
	http.SetCookie(w, &http.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 5,
	})

	http.Redirect(w, r, redirect, http.StatusFound)
}
