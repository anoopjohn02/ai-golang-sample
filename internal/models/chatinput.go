package models

type ChatInput struct {
	Question string `json:"question"`
	NewChat  bool   `json:"new_chat"`
}
