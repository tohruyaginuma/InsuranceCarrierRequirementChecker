package registry

import (
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/handler"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/store"
)

type Registry struct {
	ApplicantHandler applicantHandler
}

func NewRegistry() *Registry {
	applicantStore := store.NewApplicant()
	applicantHandler := handler.NewApplicant(applicantStore)

	return &Registry{
		ApplicantHandler: applicantHandler,
	}
}
