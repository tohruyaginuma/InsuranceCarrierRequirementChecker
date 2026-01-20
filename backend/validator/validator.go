package validator

import (
	"regexp"

	"github.com/tohruyaginuma/InsuranceCarrierRequirementChecker/model"
)

func validateEquals(value any, values []any) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}

	return false
}
func validateNotEquals(value any, values []any) bool {
	for _, v := range values {
		if value == v {
			return false
		}
	}

	return true
}

func validateRegex(value any, test string) bool {
	valueStr, ok := value.(string)
	if !ok {
		return false
	}
	matched, err := regexp.MatchString(test, valueStr)
	if err != nil {
		return false
	}
	return matched
}

func checkValidator(value any, validator model.Validator) bool {
	switch validator.Type {
	case "equals":
		return validateEquals(value, validator.Values)
	case "notEquals":
		return validateNotEquals(value, validator.Values)
	case "regex":
		return validateRegex(value, validator.Test)
	}

	return false
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

func ValidateCarrierRequirements(fieldValues []model.FieldValue, requirements []model.CarrierRequirement) (results model.ValidationResult) {
	fieldValidationResult := make([]model.FieldValidationResult, len(requirements))

	for i, r := range requirements {
		if len(r.Validator.Conditions) > 0 {
			var results = make([]bool, len(r.Validator.Conditions))
			for i, c := range r.Validator.Conditions {
				targetFieldValueForCondition := findTargetFieldValue(c.PropertyName, fieldValues)

				if targetFieldValueForCondition != nil {
					results[i] = checkValidator(targetFieldValueForCondition.Value, c.Validator)
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
			isValid = checkValidator(targetFieldValue.Value, r.Validator)
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
