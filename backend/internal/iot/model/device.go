package model

type Device struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Limit float64 `json:"limit"`
	Email string  `json:"email"`
}
