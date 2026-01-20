package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/model"
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
	slog.Debug("applicant create")

	var request createApplicantRequest
	if err := c.Bind(&request); err != nil {
		slog.Debug("applicant create bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":  "NG",
			"message": "invalid request body",
		})
	}

	applicant := model.Applicant{
		GivenName:       request.GivenName,
		Surname:         request.Surname,
		DateOfBirth:     request.DateOfBirth,
		InsuranceStatus: request.InsuranceStatus,
		PriorCarrier:    request.PriorCarrier,
		UMPD:            request.UMPD,
		Collision:       request.Collision,
	}

	applicantID := h.store.Create(c.Request().Context(), applicant)

	return c.JSON(http.StatusCreated, map[string]any{
		"result":       "OK",
		"applicant_id": applicantID,
	})
}

func (h *applicant) List(c echo.Context) error {
	slog.Debug("applicant list")

	applicants := h.store.List(c.Request().Context())

	applicantResponse := make([]listApplicantResponse, len(applicants))
	for i, a := range applicants {
		applicantResponse[i] = listApplicantResponse{
			GivenName:       a.GivenName,
			Surname:         a.Surname,
			DateOfBirth:     a.DateOfBirth,
			InsuranceStatus: a.InsuranceStatus,
			PriorCarrier:    a.PriorCarrier,
			UMPD:            a.UMPD,
			Collision:       a.Collision,
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"result":     "OK",
		"applicants": applicantResponse,
	})
}
