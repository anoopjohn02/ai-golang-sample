package main

import (
	"context"
	"log"

	"github.com/anoopjohn02/ai-golang-sample/internal/commons"
	"github.com/anoopjohn02/ai-golang-sample/internal/service"
	"github.com/anoopjohn02/ai-golang-sample/internal/web"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Starting Sample App...")
	context := Context()
	doc := service.NewDocumentService(context)
	webService := web.NewService(context, doc)
	webService.Start()
	log.Printf("Application Started...")
}

func Context() *commons.AIContext {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	ai := service.NewAIService()
	llm, _ := ai.BuildLLM()
	return &commons.AIContext{
		Context: ctx,
		LLM:     llm,
	}
}
