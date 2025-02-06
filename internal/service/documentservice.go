package service

import (
	"context"
	"log"
	"strings"

	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

type DocumentService struct {
	ctx     context.Context
	wvStore weaviate.Store
}

func NewDocumentService() *DocumentService {
	return &DocumentService{}
}

func (d *DocumentService) search(query string) (string, error) {
	docs, err := d.wvStore.SimilaritySearch(d.ctx, qr.Content, 3)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var docsContents []string
	for _, doc := range docs {
		docsContents = append(docsContents, doc.PageContent)
	}
	return strings.Join(docsContents, "\n"), nil
}
