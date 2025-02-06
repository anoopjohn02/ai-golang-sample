package service

import (
	"log"

	"github.com/tmc/langchaingo/llms/openai"
)

type AIService struct {
}

func NewAIService() *AIService {
	return &AIService{}
}

func (a *AIService) BuildLLM() (*openai.LLM, error) {
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return llm, nil
}
