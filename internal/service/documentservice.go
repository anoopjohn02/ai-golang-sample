package service

import (
	"fmt"
	"net/http"
)

type DocumentService struct {
}

func NewDocumentService() *DocumentService {
	return &DocumentService{}
}

func (d *DocumentService) search(query string) (string, error) {
	docs, err := rs.wvStore.SimilaritySearch(rs.ctx, qr.Content, 3)
	if err != nil {
		http.Error(w, fmt.Errorf("similarity search: %w", err).Error(), http.StatusInternalServerError)
		return nill, err
	}
	var docsContents []string
	for _, doc := range docs {
		docsContents = append(docsContents, doc.PageContent)
	}
}
