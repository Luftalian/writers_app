package usecase

import (
	"github.com/Luftalian/writers_app/domain"
)

type TagRepository interface {
	Store(tag domain.Tag) (domain.Tag, error)
	FindAll() (domain.Tags, error)
	FindByID(id domain.UUID) (domain.Tag, error)
}
