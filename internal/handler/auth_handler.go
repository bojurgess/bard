package handler

import (
	"fmt"
	"github.com/bojurgess/bard/internal/config"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{
		"scope":         "user-read-currently-playing",
		"state":         randomString(15),
		"response_type": "code",
		"redirect_uri":  config.AppConfig.SpotifyRedirectUri,
		"client_id":     config.AppConfig.SpotifyClientId,
	}

	q, err := mapToQueryString(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	redirect := fmt.Sprintf("https://accounts.spotify.com/authorize?%s", q)
	http.Redirect(w, r, redirect, http.StatusFound)
}

func mapToQueryString(m map[string]string) (string, error) {
	var q []string

	for key, value := range m {
		escapedKey := url.QueryEscape(key)
		escapedValue := url.QueryEscape(value)

		q = append(q, escapedKey+"="+escapedValue)
	}

	return strings.Join(q, "&"), nil
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}
