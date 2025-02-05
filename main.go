package main

import (
	"log"

	"github.com/anoopjohn02/ai-golang-sample/internal/commons"
	"github.com/anoopjohn02/ai-golang-sample/internal/web"
	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Starting Sample App...")
	context := Context()
	webService := web.NewService(context)
	webService.Start()
	log.Printf("Application Started...")
	/*
		ctx := context.Background()
		llm, err := openai.New()
		if err != nil {
			log.Fatal(err)
		}
		prompt := "What would be a good company name for a company that makes colorful socks?"
		completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(completion)
	*/
}

func Context() *commons.DeviceContext {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &commons.DeviceContext{}
}
