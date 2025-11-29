package handler

import (
	"main/internal/modules/user/dto"
	"main/internal/modules/user/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service service.UserServiceInterface
}

func NewController() *Controller {
	return &Controller{service: service.NewService()}
}

func (h *Controller) Register(e echo.Context) error {
	var register dto.RegisterDTO
	if err := e.Bind(register); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}
	err := h.service.Register(&register)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Deleted")
}
