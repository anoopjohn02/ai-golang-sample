package web

import (
	"context"
	"fmt"
	"log"

	"net/http"

	"github.com/anoopjohn02/ai-golang-sample/internal/models"
	"github.com/anoopjohn02/ai-golang-sample/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type ChatController struct {
	ser service.ChatService
}

func NewChatController() *ChatController {
	return &ChatController{ser: *service.NewChatService()}
}

func (s *ChatController) StreamChat(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	var chat models.ChatInput
	if err := c.ShouldBindJSON(&chat); err != nil {
		s.finishWithError(c, http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	content := s.ser.BuildContent(chat.Question)
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

func (req *ChatController) finishWithError(ctx *gin.Context, status int, err error) {
	var response = struct {
		Error string `json:"error"`
	}{Error: err.Error()}

	ctx.JSON(status, response)
}
