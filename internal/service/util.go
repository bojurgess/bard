package service

import (
	"encoding/base64"
	"github.com/bojurgess/bard/internal/config"
	"math/rand"
	"net/url"
	"strings"
)

var UtilService = &utilService{}

type utilService struct{}

func (s *utilService) EncodeBasicAuth() string {
	return base64.StdEncoding.EncodeToString([]byte(
		config.AppConfig.SpotifyClientId + ":" + config.AppConfig.SpotifyClientSecret))
}

func (s *utilService) MapToQueryString(m map[string]string) string {
	var q []string

	for key, value := range m {
		escapedKey := url.QueryEscape(key)
		escapedValue := url.QueryEscape(value)

		q = append(q, escapedKey+"="+escapedValue)
	}

	return strings.Join(q, "&")
}

func (s *utilService) RandomString(len int) string {
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		b[i] = byte(65 + rand.Intn(25))
	}
	return string(b)
}
