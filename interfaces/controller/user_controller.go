package controller

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
	"github.com/Luftalian/writers_app/interfaces/database"
	"github.com/Luftalian/writers_app/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(handler database.DbHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				DbHandler: handler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, u)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (controller *UserController) Show(c Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := controller.Interactor.UserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
