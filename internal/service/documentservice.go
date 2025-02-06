package service

import (
	"log"
	"strings"

	"github.com/anoopjohn02/ai-golang-sample/internal/commons"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

type DocumentService struct {
	ctx     *commons.AIContext
	wvStore weaviate.Store
}

func NewDocumentService(ctx *commons.AIContext) *DocumentService {

	log.Printf("Document service...")
	emb, err := embeddings.NewEmbedder(ctx.LLM)
	if err != nil {
		log.Fatal(err)
	}

	wvStore, _ := weaviate.New(
		weaviate.WithEmbedder(emb),
		weaviate.WithScheme("http"),
		weaviate.WithHost("localhost:9035"),
		weaviate.WithIndexName("Document"),
	)

	return &DocumentService{
		ctx:     ctx,
		wvStore: wvStore,
	}
}

func (d *DocumentService) readPdf(pdfPath string) ([]string, error) {
	log.Printf("Document service...")
	pdfConfig := pdfcpu.NewDefaultConfiguration()
	doc, err := api.ReadContextFile(pdfPath, pdfConfig, nil)
	if err != nil {
		return nil, err
	}

	var pages []string
	for i := 0; i < len(doc.PageCount); i++ {
		text, err := api.ExtractTextFile(pdfPath, nil, nil, i+1, i+1)
		if err != nil {
			return nil, err
		}
		pages = append(pages, text)
	}

	return pages, nil
}

func (d *DocumentService) add(docs []string) {
	log.Printf("Adding documents to vector db")
	// Store documents and their embeddings in weaviate
	var wvDocs []schema.Document
	for _, doc := range docs {
		wvDocs = append(wvDocs, schema.Document{PageContent: doc})
	}
	_, err := d.wvStore.AddDocuments(d.ctx.Context, wvDocs)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DocumentService) search(query string) (string, error) {
	docs, err := d.wvStore.SimilaritySearch(d.ctx.Context, query, 3)
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
