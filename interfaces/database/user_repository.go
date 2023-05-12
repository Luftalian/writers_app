package database

import (
	"github.com/Luftalian/writers_app/domain"
)

type UserRepository struct {
	DbHandler
}

func (repo *UserRepository) FindAll() (domain.Users, error) {
	users := domain.Users{}
	err := repo.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) FindByID(id domain.UUID) (domain.User, error) {
	users := domain.Users{}
	err := repo.Select(&users, "SELECT * FROM users WHERE user_id = ?", id)
	if err != nil {
		return domain.User{}, err
	}
	return users[0], nil
}

func (repo *UserRepository) Store(user domain.User) (domain.User, error) {
	_, err := repo.Exec("INSERT INTO users (user_id, name, created_at) VALUES (?, ?, ?)", user.UserID, user.Name, user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
