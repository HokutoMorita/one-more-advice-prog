package usecase

import (
	"clean_architecture_user_api/src/app/domain"
)

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
