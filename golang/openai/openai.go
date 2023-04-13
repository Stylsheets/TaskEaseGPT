package openai

import (
	"context"
	"fmt"
	"os"
)

import (
	openai "github.com/sashabaranov/go-openai"
)

var client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

type RequestEmbeddingOptions struct {
	Input []string
	Model openai.EmbeddingModel
}

func requestEmbedding(options RequestEmbeddingOptions) ([]openai.Embedding, error) {
	resp, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Input: options.Input,
		Model: options.Model,
	})

	if err != nil {
		fmt.Printf("requestEmbedding error: %v\n", err)
		return []openai.Embedding{}, err
	}

	return resp.Data, nil
}

type RequestCompletionOptions struct {
	Prompt          string
	Model           string
	Suffix          string
	MaxTokens       int
	Temperature     float32
	PresencePenalty float32
}

func requestCompletion(options RequestCompletionOptions) (string, error) {
	resp, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:           options.Model,
		Prompt:          options.Prompt,
		Suffix:          options.Suffix,
		MaxTokens:       options.MaxTokens,
		Temperature:     options.Temperature,
		PresencePenalty: options.PresencePenalty, // TODO: set lower than 0 to make it plan less
	})

	if err != nil {
		fmt.Printf("requestCompletion error: %v\n", err)
		return "", err
	}

	if resp.Choices[0].FinishReason != "stop" {
		fmt.Printf("requestCompletion error (FinishReason): %v\n", resp.Choices[0].FinishReason)
		return "", err
	}

	return resp.Choices[0].Text, nil
}

type RequestChatOptions struct {
	Messages    []openai.ChatCompletionMessage
	Model       string
	MaxTokens   int
	Temperature float32
}

func requestChat(options RequestChatOptions) (string, error) {
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Messages:    options.Messages,
		Model:       options.Model,
		MaxTokens:   options.MaxTokens,
		Temperature: options.Temperature,
	})

	if err != nil {
		fmt.Printf("requestChat error: %v\n", err)
		return "", err
	}

	if resp.Choices[0].FinishReason != "stop" {
		fmt.Printf("requestChat error (FinishReason): %v\n", resp.Choices[0].FinishReason)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
