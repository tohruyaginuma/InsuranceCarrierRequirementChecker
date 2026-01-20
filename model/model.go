package model

type Data struct {
	FieldDefinitions    []FieldDefinition    `json:"fieldDefinitions"`
	CarrierRequirements []CarrierRequirement `json:"carrierRequirements"`
	Examples            []Example            `json:"examples"`
}

type FieldDefinition struct {
	Label        string   `json:"label"`
	PropertyName string   `json:"propertyName"`
	Type         string   `json:"type"`
	Options      []Option `json:"options"`
}

type Option struct {
	Label string `json:"label"`
	Value any    `json:"value"`
}

type CarrierRequirement struct {
	PropertyName string    `json:"propertyName"`
	Validator    Validator `json:"validator"`
	Explanation  string    `json:"explanation"`
}

type Validator struct {
	Type       string      `json:"type"`
	Values     []any       `json:"values"`
	Test       string      `json:"test"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	PropertyName string    `json:"propertyName"`
	Validator    Validator `json:"validator"`
}

type Example struct {
	Description string       `json:"description"`
	Pass        bool         `json:"pass"`
	Values      []FieldValue `json:"values"`
}

type FieldValue struct {
	PropertyName string `json:"propertyName"`
	Value        any    `json:"value"`
}

type ValidationResult struct {
	IsValid      bool                    `json:"isValid"`
	FieldResults []FieldValidationResult `json:"fieldResults"`
}

type FieldValidationResult struct {
	PropertyName string `json:"propertyName"`
	IsValid      bool   `json:"isValid"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
