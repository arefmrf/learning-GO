package handler

import (
	"main/internal/modules/item/dto"
	"main/internal/modules/item/model"
	"main/internal/modules/item/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	svc service.ItemService
}

func NewItemHandler(svc service.ItemService) *ItemHandler {
	return &ItemHandler{svc}
}

func (h *ItemHandler) Create(c echo.Context) error {
	var body dto.CreateItemDTO
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	item := model.Item{
		Title: body.Title,
		Price: body.Price,
	}

	err := h.svc.Create(&item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) List(c echo.Context) error {
	items, err := h.svc.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := h.svc.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}
	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var body dto.UpdateItemDTO
	_ = c.Bind(&body)

	item, err := h.svc.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	if body.Title != "" {
		item.Title = body.Title
	}
	if body.Price != 0 {
		item.Price = body.Price
	}

	err = h.svc.Update(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.svc.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}
	return c.JSON(http.StatusOK, "Deleted")
}
