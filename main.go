package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/model"
	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/validator"
)

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

func findTargetFieldValue(propertyName string, fieldValues []model.FieldValue) (fieldValue *model.FieldValue) {
	for _, v := range fieldValues {
		if v.PropertyName == propertyName {
			return &v
		}
	}
	return nil
}

func everyFieldValidationResult(results []model.FieldValidationResult) (result bool) {
	for _, r := range results {
		if !r.IsValid {
			return false
		}
	}

	return true
}

func validateCarrierRequirements(fieldValues []model.FieldValue, requirements []model.CarrierRequirement) (results model.ValidationResult) {
	fieldValidationResult := make([]model.FieldValidationResult, len(requirements))

	for i, r := range requirements {
		if len(r.Validator.Conditions) > 0 {
			var results = make([]bool, len(r.Validator.Conditions))
			for i, c := range r.Validator.Conditions {
				targetFieldValueForCondition := findTargetFieldValue(c.PropertyName, fieldValues)

				if targetFieldValueForCondition != nil {
					results[i] = validator.CheckValidator(targetFieldValueForCondition.Value, c.Validator)
				}
			}

			var result bool
			for _, r := range results {
				if !r {
					result = false
					break
				}

				result = true
			}
			if !result {
				fieldValidationResult[i] = model.FieldValidationResult{
					PropertyName: r.PropertyName,
					IsValid:      true,
				}
				continue
			}
		}

		targetFieldValue := findTargetFieldValue(r.PropertyName, fieldValues)

		var isValid bool
		if targetFieldValue != nil {
			isValid = validator.CheckValidator(targetFieldValue.Value, r.Validator)
		}
		var errorMessage string
		if !isValid {
			errorMessage = r.Explanation
		}

		fieldValidationResult[i] = model.FieldValidationResult{
			PropertyName: r.PropertyName,
			IsValid:      isValid,
			ErrorMessage: errorMessage,
		}
	}

	isValid := everyFieldValidationResult(fieldValidationResult)

	return model.ValidationResult{
		IsValid:      isValid,
		FieldResults: fieldValidationResult,
	}
}

func main() {
	data := unmarshalData()

	for i, e := range data.Examples {
		fmt.Println(strings.Repeat("=", 60))
		fmt.Printf("Test %d: %s\n", i+1, e.Description)
		fmt.Printf("Expected to pass: %v\n", e.Pass)
		fmt.Println(strings.Repeat("=", 60))

		result := validateCarrierRequirements(e.Values, data.CarrierRequirements)

		var resultText string
		if result.IsValid {
			resultText = "✓ VALID"
		} else {
			resultText = "✗ INVALID"
		}

		var matchReultText string
		if result.IsValid == e.Pass {
			matchReultText = "✓"
		} else {
			matchReultText = "✗"
		}

		fmt.Printf("Result: %v\n", resultText)
		fmt.Printf("Match expected: %v\n", matchReultText)

		if len(result.FieldResults) > 0 {
			fmt.Println("Field Results:")
			for _, fr := range result.FieldResults {
				if !fr.IsValid {
					fmt.Printf("  ✗ %v: %v\n", fr.PropertyName, fr.ErrorMessage)
				}
			}
		}
	}
}
