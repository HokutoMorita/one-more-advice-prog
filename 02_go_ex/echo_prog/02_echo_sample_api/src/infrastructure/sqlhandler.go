package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"echo_sample_api/interfaces/database"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "root:one_more_advice_prog@tcp(127.0.0.1:4307)/one_more_advice_prog?collation=utf8mb4_unicode_ci&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) FindById(obj interface{}, id int) {
	handler.db.First(obj, id)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}

func (handler *SqlHandler) UpdateAll(obj interface{}) {
	handler.db.Save(obj)
}
