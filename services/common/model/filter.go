package model

type Filter struct {
	Name string `json:"name"`

	Operator *string `json:"operator"`

	Values []string `json:"values"`
}
