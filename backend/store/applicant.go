package store

import (
	"context"

	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/model"
)

type applicant struct {
	applicants []model.Applicant
	nextID     int64
}

func NewApplicant() *applicant {
	return &applicant{
		applicants: []model.Applicant{},
		nextID:     1,
	}
}

func (s *applicant) Create(ctx context.Context, applicant model.Applicant) int64 {
	applicant.ID = s.nextID
	s.nextID++
	s.applicants = append(s.applicants, applicant)
	return applicant.ID
}

func (s *applicant) List(ctx context.Context) []model.Applicant {
	return s.applicants
}
