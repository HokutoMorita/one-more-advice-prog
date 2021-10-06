package database

import (
	"echo_sample_api/domain"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u domain.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []domain.User {
	user := []domain.User{}
	db.FindAll(&user)
	return user
}

func (db *UserRepository) Delete(id string) {
	user := []domain.User{}
	db.DeleteById(&user, id)
}

func (db *UserRepository) Update(u domain.User) {
	db.UpdateAll(&u)
}
