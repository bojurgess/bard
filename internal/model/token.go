package model

import "time"

type DatabaseTokens struct {
	ID           int       `db:"id"`
	UserID       string    `db:"user_id"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	ExpiresAt    time.Time `db:"expires_at"`
}

type OAuthTokens struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
}

func OAuthToDatabaseTokens(ot *OAuthTokens, userId string) *DatabaseTokens {
	return &DatabaseTokens{
		UserID:       userId,
		AccessToken:  ot.AccessToken,
		RefreshToken: ot.RefreshToken,
		ExpiresAt:    time.Now().Add(ot.ExpiresIn),
	}
}
