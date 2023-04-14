package openai

import (
	"TaskEaseGPT/config"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var client = openai.NewClient(config.GetOpenAICredentials().ApiKey)

type RequestChatOptions struct {
	Messages    []openai.ChatCompletionMessage
	Model       string
	MaxTokens   int
	Temperature float32
}

func RequestChat(options RequestChatOptions) (string, error) {
	if options.MaxTokens == 0 {
		options.MaxTokens, _ = CalculateTokens(options.Messages[len(options.Messages)-1].Content)
	}
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

type RequestEmbeddingOptions struct {
	Input []string
	Model openai.EmbeddingModel
}

func RequestEmbedding(options RequestEmbeddingOptions) ([]openai.Embedding, error) {
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

func RequestCompletion(options RequestCompletionOptions) (string, error) {
	resp, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:           options.Model,
		Prompt:          options.Prompt,
		Suffix:          options.Suffix,
		MaxTokens:       options.MaxTokens,
		Temperature:     options.Temperature,
		PresencePenalty: options.PresencePenalty, // TODO: set lower than 0 to make it plan less
	})

	if err != nil {
		return "", err
	}

	if resp.Choices[0].FinishReason != "stop" {
		err = fmt.Errorf("requestCompletion error (FinishReason): %v", resp.Choices[0].FinishReason)
		return "", err
	}

	return resp.Choices[0].Text, nil
}
