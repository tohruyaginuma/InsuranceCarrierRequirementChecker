package registry

import (
	"encoding/json"
	"os"

	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/handler"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/model"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/store"
)

type Registry struct {
	ApplicantHandler applicantHandler
}

func unmarshalData() model.Data {
	file, err := os.ReadFile("data/data.json")
	if err != nil {
		panic(err)
	}

	var data model.Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	return data
}

func NewRegistry() *Registry {
	jsonData := unmarshalData()

	applicantStore := store.NewApplicant()
	applicantHandler := handler.NewApplicant(applicantStore, jsonData.CarrierRequirements)

	return &Registry{
		ApplicantHandler: applicantHandler,
	}
}
