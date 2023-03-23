package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserLikeRepository interface {
	Store(userLike domain.UserLike) (domain.UserLike, error)
	FindAll() (domain.UserLikes, error)
	FindByUserID(userID uuid.UUID) (domain.UserLike, error)
	FindByTextID(textID uuid.UUID) (domain.UserLike, error)
	Update(userLike domain.UserLike) error
	Delete(u domain.UserLike) error
	DeleteByUserID(userID uuid.UUID) error
	DeleteByTextID(textID uuid.UUID) error
}
