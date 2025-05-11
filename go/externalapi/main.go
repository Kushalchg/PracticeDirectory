package main

import (
	// "context"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// ctx := context.Background()
	// GoogleApiCall(ctx)
	// MistralApiCall()
	OpenRouter()

}
