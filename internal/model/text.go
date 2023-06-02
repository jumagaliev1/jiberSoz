package model

type Text struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	Link    string `json:"link"`
}
