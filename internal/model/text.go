package model

type Text struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	Link    string `json:"link"`
}

func (t Text) ToTextResponse() TextResponse {
	return TextResponse{
		ID:      t.ID,
		Message: t.Message,
		Link:    t.Link,
	}
}

type TextRequest struct {
	Message string `json:"message"`
}

func (t TextRequest) ToText() Text {
	return Text{
		Message: t.Message,
	}
}

type TextResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
	Link    string `json:"link"`
}
