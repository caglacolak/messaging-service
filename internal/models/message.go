package models

type Message struct {
	ID        int     `json:"id" example:"1"`
	Content   string  `json:"content" example:"Hello, world!"`
	Recipient string  `json:"recipient" example:"+905383311137"`
	Status    string  `json:"status" example:"sent"`
	SentAt    *string `json:"sent_at"`
}
