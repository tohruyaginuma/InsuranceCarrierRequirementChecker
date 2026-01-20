package handler

import (
	"context"

	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/model"
)

type applicantStore interface {
	Create(ctx context.Context, applicant model.Applicant) int64
	List(ctx context.Context) []model.Applicant
}
