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

func CheckValidator(value any, validator model.Validator) bool {
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
