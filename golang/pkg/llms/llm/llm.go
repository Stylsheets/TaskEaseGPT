package llm

import (
	"TaskEaseGPT/pkg/llms/openai"
	"fmt"
	openai2 "github.com/sashabaranov/go-openai"
)

type LLM interface {
	RequestChat(options RequestChatOptions) (string, error)
	RequestEmbedding(options RequestEmbeddingOptions) ([]Embedding, error)
	RequestCompletion(options RequestCompletionOptions) (string, error)
}

// Define common types for input and output
type RequestChatOptions struct {
	Messages    []ChatCompletionMessage
	Model       string
	MaxTokens   int
	Temperature float32
}

type ChatCompletionMessage struct {
	Role    string
	Content string
}

type RequestEmbeddingOptions struct {
	Input []string
	Model string
}

type Embedding struct {
	ID    string
	Value []float32
}

type RequestCompletionOptions struct {
	Prompt          string
	Model           string
	Suffix          string
	MaxTokens       int
	Temperature     float32
	PresencePenalty float32
}

type OpenAIAdapter struct {
	Client *openai.Client
}

func NewOpenAIAdapter(apiKey string) *OpenAIAdapter {
	return &OpenAIAdapter{
		Client: openai.NewClient(apiKey),
	}
}

func (oa *OpenAIAdapter) RequestChat(options RequestChatOptions) (string, error) {
	messages := make([]openai2.ChatCompletionMessage, len(options.Messages))
	for i, m := range options.Messages {
		messages[i] = openai2.ChatCompletionMessage{
			Role:    m.Role,
			Content: m.Content,
		}
	}

	return oa.Client.RequestChat(openai.RequestChatOptions{
		Messages:    messages,
		Model:       options.Model,
		MaxTokens:   options.MaxTokens,
		Temperature: options.Temperature,
	})
}

func (oa *OpenAIAdapter) RequestEmbedding(options RequestEmbeddingOptions) ([]Embedding, error) {
	embeddings, err := oa.Client.RequestEmbedding(openai.RequestEmbeddingOptions{
		Input: options.Input,
		Model: openai2.AdaEmbeddingV2,
	})
	if err != nil {
		return nil, err
	}

	results := make([]Embedding, len(embeddings))
	for i, e := range embeddings {
		results[i] = Embedding{
			ID:    fmt.Sprint(e.Index),
			Value: e.Embedding,
		}
	}

	return results, nil
}

func (oa *OpenAIAdapter) RequestCompletion(options RequestCompletionOptions) (string, error) {
	return oa.Client.RequestCompletion(openai.RequestCompletionOptions(options))
}
