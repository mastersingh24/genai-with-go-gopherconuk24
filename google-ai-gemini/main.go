package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// import standard libraries
	// Import the GenerativeAI package for Go
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()
	// Access your API key as an environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_GENAI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.9)
	model.SetTopP(0.5)
	model.SetTopK(20)
	//model.SetMaxOutputTokens(100)
	//model.SystemInstruction = genai.NewUserContent(genai.Text("You are Yoda from Star Wars."))
	model.ResponseMIMEType = "application/json"

	iter := model.GenerateContentStream(ctx, genai.Text("Write a story about a magic backpack."))
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		printResponse(resp)
	}

}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
