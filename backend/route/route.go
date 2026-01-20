package route

import (
	"github.com/labstack/echo/v4"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/registry"
)

func SetRoute(e *echo.Echo, r *registry.Registry) {
	const version = "v1"

	applicantGroup := e.Group(version + "/applicants")

	applicantGroup.POST("/", r.ApplicantHandler.Create)
	applicantGroup.GET("/", r.ApplicantHandler.List)
}
