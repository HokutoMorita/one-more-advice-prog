package main

import (
	"echo_sample_api/domain"
	"echo_sample_api/infrastructure"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var (
	db *gorm.DB
	err error
	dsn = "root:one_more_advice_prog@tcp(127.0.0.1:4307)/one_more_advice_prog?collation=utf8mb4_unicode_ci&parseTime=true"
)

func main() {
	dbinit()
	infrastructure.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}

func dbinit() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	db.Migrator().CreateTable(domain.User{})
	fmt.Println("DB Migration成功")
}
