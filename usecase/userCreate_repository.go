package usecase

import (
	"github.com/Luftalian/writers_app/domain"
)

type UserCreateRepository interface {
	Store(userCreate domain.UserCreate) (domain.UserCreate, error)
	FindAll() (domain.UserCreates, error)
	FindByUserID(userID domain.UUID) (domain.Texts, error)
	FindByTextID(textID domain.UUID) (domain.UserCreate, error)
	Update(userCreate domain.UserCreate) error
	Delete(u domain.UserCreate) error
	DeleteByUserID(userID domain.UUID) error
	DeleteByTextID(textID domain.UUID) error
}
