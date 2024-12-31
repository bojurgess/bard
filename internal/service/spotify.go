package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/bojurgess/bard/internal/config"
	"github.com/bojurgess/bard/internal/model"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

var SpotifyService = &spotifyService{}

type spotifyService struct{}

func (s *spotifyService) GenerateAuthUrl() (string, string) {
	state := randomString(15)
	params := map[string]string{
		"scope":         "user-read-currently-playing",
		"state":         state,
		"response_type": "code",
		"redirect_uri":  config.AppConfig.SpotifyRedirectUri,
		"client_id":     config.AppConfig.SpotifyClientId,
	}

	q := mapToQueryString(params)
	return string("https://accounts.spotify.com/authorize?" + q), state
}

func (s *spotifyService) RequestAccessToken(code string) (*model.OAuthTokens, error) {
	var tokens model.OAuthTokens

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + encodeBasicAuth(),
	}

	body := mapToQueryString(map[string]string{
		"code":         code,
		"redirect_uri": config.AppConfig.SpotifyRedirectUri,
		"grant_type":   "authorization_code",
	})

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (s *spotifyService) RefreshAccessToken(refreshToken string) (*model.OAuthTokens, error) {
	var tokens model.OAuthTokens

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Bearer " + refreshToken,
	}

	body := mapToQueryString(map[string]string{
		"client_id":     config.AppConfig.SpotifyClientId,
		"client_secret": config.AppConfig.SpotifyClientSecret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	})

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (*spotifyService) Me(accessToken string) (*model.User, error) {
	var user model.User
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func encodeBasicAuth() string {
	return base64.StdEncoding.EncodeToString([]byte(
		config.AppConfig.SpotifyClientId + ":" + config.AppConfig.SpotifyClientSecret))
}

func mapToQueryString(m map[string]string) string {
	var q []string

	for key, value := range m {
		escapedKey := url.QueryEscape(key)
		escapedValue := url.QueryEscape(value)

		q = append(q, escapedKey+"="+escapedValue)
	}

	return strings.Join(q, "&")
}

func randomString(len int) string {
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		b[i] = byte(65 + rand.Intn(25))
	}
	return string(b)
}
