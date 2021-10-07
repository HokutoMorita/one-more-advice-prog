package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:one_more_advice_prog@tcp(127.0.0.1:4307)/one_more_advice_prog?collation=utf8mb4_unicode_ci&parseTime=true")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
