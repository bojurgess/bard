package model

type User struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"display_name" db:"display_name"`
}

type UserWithTokens struct {
	User   User           `json:"user"`
	Tokens DatabaseTokens `json:"tokens"`
}
