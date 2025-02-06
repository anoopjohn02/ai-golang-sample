package commons

import (
	"context"

	"github.com/tmc/langchaingo/llms/openai"
)

type AIContext struct {
	Context context.Context
	LLM     *openai.LLM
}
