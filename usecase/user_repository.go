package usecase

import (
	"github.com/Luftalian/writers_app/domain"
)

type UserRepository interface {
	Store(user domain.User) (domain.User, error)
	FindAll() (domain.Users, error)
	FindByID(id domain.UUID) (domain.User, error)
}
