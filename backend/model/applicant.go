package model

type Applicant struct {
	ID              int64
	GivenName       string
	Surname         string
	DateOfBirth     string
	InsuranceStatus string
	PriorCarrier    string
	UMPD            *int
	Collision       *int
}
