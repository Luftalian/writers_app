package controller

import "github.com/Luftalian/writers_app/domain"

type UUIDHandler interface {
	New() domain.UUID
	Parse(s string) (domain.UUID, error)
}
