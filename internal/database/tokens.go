package database

import "github.com/bojurgess/bard/internal/model"

var TokenService = &tokenService{}

type tokenService struct{}

func (t *tokenService) Create(tokens *model.DatabaseTokens) error {
	_, err := db.NamedExec(`INSERT INTO tokens (user_id, access_token, refresh_token, expires_at) VALUES (:user_id, :access_token, :refresh_token, :expires_at) ON CONFLICT DO UPDATE SET access_token = excluded.access_token, expires_at = excluded.expires_at`, tokens)
	return err
}

func (t *tokenService) Find(userID string) (*model.DatabaseTokens, error) {
	var token model.DatabaseTokens
	err := db.Get(&token, "SELECT * FROM tokens WHERE user_id = $1", userID)
	return &token, err
}

func (t *tokenService) Update(tokens *model.DatabaseTokens) error {
	_, err := db.NamedExec(`UPDATE tokens SET (access_token, expires_at) = (:access_token, :expires_at) WHERE user_id = :user_id`, tokens)
	return err
}

func (t *tokenService) Delete(userID string) error {
	_, err := db.NamedExec("DELETE FROM tokens WHERE user_id = $1", userID)
	return err
}
