package service

import (
	"fmt"

	"github.com/anoopjohn02/ai-golang-sample/internal/commons"
	"github.com/tmc/langchaingo/llms"
)

type ChatService struct {
	ctx *commons.AIContext
	doc *DocumentService
}

func NewChatService(ctx *commons.AIContext, doc *DocumentService) *ChatService {
	return &ChatService{
		ctx: ctx,
		doc: doc,
	}
}

func (s *ChatService) BuildContent(query string) []llms.MessageContent {
	docsContents, _ := s.doc.search(query)
	context := fmt.Sprintf(systemMessage, docsContents)
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, context),
		llms.TextParts(llms.ChatMessageTypeHuman, query),
	}
	return content
}

const systemMessage = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Context:
%s
`
