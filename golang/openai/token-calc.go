package openai

import (
	"github.com/tiktoken-go/tokenizer"
)

func CalculateTokens(input string) (int, error) {
	enc, err := tokenizer.Get(tokenizer.Cl100kBase)
	if err != nil {
		return 0, err
	}

	ids, _, err := enc.Encode(input)
	return len(ids), err
}
