package web

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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

	c.Writer.Flush()
	for i := 0; i < 10; i++ {
		_, err := fmt.Fprintf(c.Writer, "data: Message %d ", i)
		if err != nil {
			return
		}
		c.Writer.Flush()
		time.Sleep(1 * time.Second)
	}
}
