package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type applicant struct {
	store applicantStore
}

func NewApplicant(store applicantStore) *applicant {
	return &applicant{
		store: store,
	}
}

func (h *applicant) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
	})
}
func (h *applicant) List(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": "OK",
	})
}
