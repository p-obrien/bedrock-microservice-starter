package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
	"github.com/aws/aws-sdk-go/aws"
)

type Question struct {
	Text string `json:"text"`
}

// Sample API Key
var apiKey string = "test-api-key"
var brc *bedrockruntime.Client

const modelID = "anthropic.claude-3-sonnet-20240229-v1:0"

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	brc = bedrockruntime.NewFromConfig(cfg)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == apiKey, nil
	}))

	e.POST("/converse", converseHandler)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func converseHandler(c echo.Context) error {
	// Initialize the Bedrock ConverseInput with the model ID
	converseInput := &bedrockruntime.ConverseInput{
		ModelId: aws.String(modelID),
	}

	// Get the user input from the form value
	input := c.FormValue("message")

	// Create the user's message
	userMsg := types.Message{
		Role: types.ConversationRoleUser,
		Content: []types.ContentBlock{
			&types.ContentBlockMemberText{
				Value: input,
			},
		},
	}

	// Append the user's message to the conversation input
	converseInput.Messages = append(converseInput.Messages, userMsg)

	// Call the Bedrock Converse API
	output, err := brc.Converse(context.Background(), converseInput)
	if err != nil {
		log.Fatal("Error calling Converse API:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Extract the response from the assistant
	response, ok := output.Output.(*types.ConverseOutputMemberMessage)
	if !ok {
		return c.String(http.StatusInternalServerError, "Failed to parse response")
	}

	responseContentBlock := response.Value.Content[0]
	text, ok := responseContentBlock.(*types.ContentBlockMemberText)
	if !ok {
		return c.String(http.StatusInternalServerError, "Failed to parse response content")
	}

	// Create the assistant's message
	assistantMsg := types.Message{
		Role:    types.ConversationRoleAssistant,
		Content: response.Value.Content,
	}

	// Append the assistant's message to the conversation input
	converseInput.Messages = append(converseInput.Messages, assistantMsg)

	// Return the assistant's response to the client
	return c.JSON(http.StatusOK, text.Value)
}
