package controllers

import (
	"fmt"
	"echo_sample_api/domain"
	"echo_sample_api/interfaces/database"
	"echo_sample_api/usecase"

	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := domain.User{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []domain.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}

func (controller *UserController) Update(c echo.Context) {
	u := domain.User{}
	c.Bind(&u)
	fmt.Println("取得した値")
	fmt.Println(u)
	controller.Interactor.Update(u)
}
