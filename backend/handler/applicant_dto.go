package handler

type listApplicantResponse struct {
	ID              int64  `json:"id"`
	GivenName       string `json:"given_name"`
	Surname         string `json:"surname"`
	DateOfBirth     string `json:"date_of_birth"`
	InsuranceStatus string `json:"insurance_status"`
	PriorCarrier    string `json:"prior_carrier"`
	UMPD            *int   `json:"umpd"`
	Collision       *int   `json:"collision"`
}

type createApplicantRequest struct {
	GivenName       string `json:"given_name"`
	Surname         string `json:"surname"`
	DateOfBirth     string `json:"date_of_birth"`
	InsuranceStatus string `json:"insurance_status"`
	PriorCarrier    string `json:"prior_carrier"`
	UMPD            *int   `json:"umpd"`
	Collision       *int   `json:"collision"`
}
