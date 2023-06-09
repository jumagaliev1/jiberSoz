package handler

import (
	"github.com/jumagaliev1/jiberSoz/internal/model"
	"github.com/jumagaliev1/jiberSoz/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TextHandler struct {
	service *service.Service
}

func NewTextHandler(service *service.Service) *TextHandler {
	return &TextHandler{
		service: service,
	}
}

func (h *TextHandler) Create(c echo.Context) error {
	var body model.TextRequest

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	text, err := h.service.Text.Create(c.Request().Context(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{"link:": text.Link})

}

func (h *TextHandler) GetByLink(c echo.Context) error {
	link := c.Param("link")

	text, err := h.service.Text.GetByLink(c.Request().Context(), link)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, text)
}
