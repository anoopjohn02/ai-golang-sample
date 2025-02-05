package web

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type ChatController struct {
}

func NewChatController() *ChatController {
	return &ChatController{}
}

func (s *ChatController) StreamChat(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	prompt := "Tell me a story about a wise old owl."

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a story teller."),
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	}
	completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Fprintf(c.Writer, string(chunk))
		c.Writer.Flush()
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}
	_ = completion

}
