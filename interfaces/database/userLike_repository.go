package database

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserLikeRepository struct {
	DbHandler
}

func (repo *UserLikeRepository) FindAll() (domain.UserLikes, error) {
	userLikes := domain.UserLikes{}
	err := repo.Select(&userLikes, "SELECT * FROM userLikes")
	if err != nil {
		return nil, err
	}
	return userLikes, nil
}

func (repo *UserLikeRepository) FindByUserID(userID uuid.UUID) (domain.UserLike, error) {
	userLikes := domain.UserLikes{}
	err := repo.Select(&userLikes, "SELECT * FROM userLikes WHERE user_id = ?", userID)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLikes[0], nil
}

func (repo *UserLikeRepository) FindByTextID(textID uuid.UUID) (domain.UserLike, error) {
	userLikes := domain.UserLikes{}
	err := repo.Select(&userLikes, "SELECT * FROM userLikes WHERE text_id = ?", textID)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLikes[0], nil
}

func (repo *UserLikeRepository) Store(userLike domain.UserLike) (domain.UserLike, error) {
	_, err := repo.Exec("INSERT INTO userLikes (user_id, text_id) VALUES (?, ?)", userLike.UserID, userLike.TextID)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLike, nil
}

func (repo *UserLikeRepository) Update(userLike domain.UserLike) error {
	_, err := repo.Exec("UPDATE userLikes SET user_id = ?, text_id = ? WHERE user_id = ? AND text_id = ?", userLike.UserID, userLike.TextID, userLike.UserID, userLike.TextID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserLikeRepository) Delete(u domain.UserLike) error {
	_, err := repo.Exec("DELETE FROM userLikes WHERE user_id = ? AND text_id = ?", u.UserID, u.TextID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserLikeRepository) DeleteByUserID(userID uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM userLikes WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserLikeRepository) DeleteByTextID(textID uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM userLikes WHERE text_id = ?", textID)
	if err != nil {
		return err
	}
	return nil
}
