package usecase

import (
	"github.com/Luftalian/writers_app/domain"
)

type TextRepository interface {
	Store(text domain.Text) (domain.Text, error)
	FindAll() (domain.Texts, error)
	FindByID(id domain.UUID) (domain.Text, error)
	Update(text domain.Text) error
	DeleteByID(id domain.UUID) error
}
