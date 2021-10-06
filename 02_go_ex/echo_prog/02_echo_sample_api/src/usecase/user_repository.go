package usecase

import (
	"echo_sample_api/domain"
)

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	SelectById(id string) domain.User
	Delete(id string)
	Update(domain.User)
}
