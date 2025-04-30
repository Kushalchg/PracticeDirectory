package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		fmt.Println("Error occured in creating client::", err)
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		SystemInstruction: genai.NewContentFromText(
			`You are nepali to english translator and you convert everything 
      I given in nepali into english, content may contain the person anme location and other stuff from country nepal, you just convert into english`,
			genai.RoleUser),
	}

	res, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(""),
		config,
	)

	f, err := os.OpenFile("./result.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error while opeinig file::", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	value, err := w.WriteString(res.Text())
	fmt.Println(value)

	if err != nil {
		fmt.Println("error while writing into file::", err)
	}

	err = w.Flush()
	if err != nil {
		fmt.Println("error while flushing buffer:", err)
	}
}
