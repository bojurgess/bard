package database

import "github.com/bojurgess/bard/internal/model"

var UserService = &userService{}

type userService struct{}

func (u *userService) Create(user *model.User) error {
	_, err := db.NamedExec(`INSERT INTO users (id, display_name) VALUES (:id, :display_name) ON CONFLICT DO NOTHING`, user)
	return err
}

func (u *userService) Find(id string) (*model.User, error) {
	var user model.User
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	return &user, err
}

func (u *userService) FindWithTokens(id string) (*model.UserWithTokens, error) {
	var user model.UserWithTokens
	err := db.Get(&user, "SELECT * FROM users INNER JOIN main.tokens t on users.id = t.user_id WHERE users.id = $1", id)
	return &user, err
}

func (u *userService) Update(user *model.User) error {
	_, err := db.NamedExec(`UPDATE users SET display_name = :display_name WHERE id = :id`, user)
	return err
}

func (u *userService) Delete(id string) error {
	_, err := db.NamedExec(`DELETE FROM users WHERE id = $1`, id)
	return err
}

func (u *userService) Exists(id string) bool {
	var count int32
	err := db.Get(&count, `SELECT COUNT(*) FROM users WHERE id = $1`, id)
	if err != nil {
		return false
	}

	return count == 1
}
