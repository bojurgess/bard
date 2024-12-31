package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bojurgess/bard/internal/config"
	"github.com/bojurgess/bard/internal/model"
	"net/http"
)

var SpotifyService = &spotifyService{}

type spotifyService struct{}

func (s *spotifyService) GenerateAuthUrl() (string, string) {
	state := UtilService.RandomString(15)
	params := map[string]string{
		"scope":         "user-read-currently-playing",
		"state":         state,
		"response_type": "code",
		"redirect_uri":  config.AppConfig.SpotifyRedirectUri,
		"client_id":     config.AppConfig.SpotifyClientId,
	}

	q := UtilService.MapToQueryString(params)
	return string("https://accounts.spotify.com/authorize?" + q), state
}

func (s *spotifyService) RequestAccessToken(code string) (*model.OAuthTokens, error) {
	var tokens model.OAuthTokens

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic " + UtilService.EncodeBasicAuth(),
	}

	body := UtilService.MapToQueryString(map[string]string{
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
		"Authorization": "Basic " + UtilService.EncodeBasicAuth(),
	}

	body := UtilService.MapToQueryString(map[string]string{
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
		return nil, errors.New(resp.Status)
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

func (s *spotifyService) GetCurrentlyPlaying(accessToken string) (*model.SpotifyCurrentlyPlaying, error) {
	var currentlyPlaying model.SpotifyCurrentlyPlaying

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/currently-playing", nil)
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
		var errorResponse model.SpotifyErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(errorResponse.Error.Message)
	}

	err = json.NewDecoder(resp.Body).Decode(&currentlyPlaying)
	if err != nil {
		return nil, err
	}

	return &currentlyPlaying, nil
}
