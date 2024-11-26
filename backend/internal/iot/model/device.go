package model

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data []Data `json:"data"`
}
