package database

type SqlHandler interface {
	Create(object interface{})
	FindAll(object interface{})
	FindById(object interface{}, id string)
	DeleteById(object interface{}, id string)
	UpdateAll(object interface{})
}
