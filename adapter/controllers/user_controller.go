package controllers

import (
	"github.com/kazu1029/go-clean-arch/adapter/gateway"
	"github.com/kazu1029/go-clean-arch/adapter/interfaces"
	"github.com/kazu1029/go-clean-arch/domain"
	"github.com/kazu1029/go-clean-arch/usecase"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(conn *gorm.DB, logger interfaces.Logger) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &gateway.UserRepository{
				Conn: conn,
			},
			Logger: logger,
		},
	}
}

func (controller *UserController) Create(c interfaces.Context) {
	type (
		Request struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		Response struct {
			UserID int `json:"user_id"`
		}
	)

	// use PostForm when Content-Type is application/x-www-form-urlencoded
	user := domain.User{Name: c.PostForm("name"), Email: c.PostForm("email")}

	id, err := controller.Interactor.Add(user)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "user_controller: cannot add user"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{UserID: id}
	c.JSON(201, res)
}
