package model

import "time"

type Text struct {
	ID        uint      `json:"id"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func (t Text) ToTextResponse() TextResponse {
	return TextResponse{
		ID:        t.ID,
		Link:      t.Link,
		CreatedAt: t.CreatedAt,
		ExpiresAt: t.ExpiresAt,
	}
}

type TextRequest struct {
	Message string `json:"message"`
	Day     int    `json:"day"`
}

func (t TextRequest) ToText() Text {
	return Text{}
}

type TextResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}
