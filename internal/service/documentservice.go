package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anoopjohn02/ai-golang-sample/internal/commons"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

type DocumentService struct {
	ctx     *commons.AIContext
	wvStore weaviate.Store
}

func NewDocumentService(ctx *commons.AIContext) *DocumentService {

	log.Printf("Document service...")

	embeddingsClient, err := openai.New(openai.WithModel("text-embedding-ada-002"))
	emb, err := embeddings.NewEmbedder(embeddingsClient)
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

func (d *DocumentService) ReadFile(filePath string) ([]string, error) {
	log.Printf("Reading documents")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a slice to store the lines
	var lines []string

	// Read line by line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}

func (d *DocumentService) AddDocs(docs []string) {
	log.Printf("Adding documents to vector db")
	// Store documents and their embeddings in weaviate
	var wvDocs []schema.Document
	for _, doc := range docs {
		if strings.TrimSpace(doc) != "" {
			wvDocs = append(wvDocs, schema.Document{PageContent: doc})
		}
	}

	log.Printf("Document batch size: %d", len(wvDocs))
	batchSize := 5 // Experiment with smaller batch sizes
	for i := 0; i < len(wvDocs); i += batchSize {
		end := i + batchSize
		if end > len(wvDocs) {
			end = len(wvDocs)
		}
		log.Printf("Processing: %d - %d", i, end)
		_, err := d.wvStore.AddDocuments(d.ctx.Context, wvDocs[i:end])
		if err != nil {
			log.Fatal(err)
		}
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
