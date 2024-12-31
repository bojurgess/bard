package handler

import (
	"github.com/bojurgess/bard/internal/database"
	"github.com/bojurgess/bard/internal/model"
	"github.com/bojurgess/bard/internal/service"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	code := q.Get("code")
	errorMsg := q.Get("error")
	state := q.Get("state")

	storedState, err := r.Cookie("state")
	if err != nil || storedState == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if state != storedState.Value || errorMsg != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := service.SpotifyService.RequestAccessToken(code)
	if err != nil || tokens == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := service.SpotifyService.Me(tokens.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = database.UserService.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbTokens := model.OAuthToDatabaseTokens(tokens, user.ID)
	err = database.TokenService.Create(dbTokens)

	_, err = w.Write([]byte("Successfully authenticated! You can now close this tab."))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
