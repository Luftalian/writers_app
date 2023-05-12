package usecase

import (
	"github.com/Luftalian/writers_app/domain"
)

type UserLikeRepository interface {
	Store(userLike domain.UserLike) (domain.UserLike, error)
	FindAll() (domain.UserLikes, error)
	CheckLikeUser(domain.UserLike) (bool, error)
	FindByUserID(userID domain.UUID) (domain.Texts, error)
	FindByTextID(textID domain.UUID) (domain.UserLike, error)
	Update(userLike domain.UserLike) error
	Delete(u domain.UserLike) error
	DeleteByUserID(userID domain.UUID) error
	DeleteByTextID(textID domain.UUID) error
}

const (
	DatabaseError = iota
	NoRowsError
	ExitData
)
