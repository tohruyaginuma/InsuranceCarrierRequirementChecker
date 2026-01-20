package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/model"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/validator"
)

type applicant struct {
	store        applicantStore
	requirements []model.CarrierRequirement
}

func NewApplicant(store applicantStore, requirements []model.CarrierRequirement) *applicant {
	return &applicant{
		store:        store,
		requirements: requirements,
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

	fieldValues := []model.FieldValue{
		{PropertyName: "ApplicantGivenName", Value: request.GivenName},
		{PropertyName: "ApplicantSurname", Value: request.Surname},
		{PropertyName: "ApplicantDOB", Value: request.DateOfBirth},
		{PropertyName: "InsuranceStatus", Value: request.InsuranceStatus},
		{PropertyName: "PriorCarrier", Value: request.PriorCarrier},
		{PropertyName: "UMPD", Value: request.UMPD},
		{PropertyName: "Collision", Value: request.Collision},
	}

	result := validator.ValidateCarrierRequirements(fieldValues, h.requirements)
	if !result.IsValid {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"result":       "NG",
			"fieldResults": result.FieldResults,
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
			ID:              a.ID,
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
