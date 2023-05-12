package infrastructure

import (
	"github.com/Luftalian/writers_app/domain"
	"github.com/google/uuid"
)

type UUIDHandler struct{}

func (u *UUIDHandler) New() domain.UUID {
	return uuid.New()
}

func (u *UUIDHandler) Parse(s string) (domain.UUID, error) {
	return uuid.Parse(s)
}
