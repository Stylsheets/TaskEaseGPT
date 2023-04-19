package google

import (
	"context"
	"errors"
	"fmt"

	"github.com/rocketlaunchr/google-search"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

type Google struct {
	JsonKeyFilePath string
	CustomCX        string
}

type SearchResult struct {
	Rank        int
	URL         string
	Title       string
	Description string
}

func (g Google) searchAnonymous(query string) ([]SearchResult, error) {
	results, err := googlesearch.Search(nil, query)

	searchResults := make([]SearchResult, len(results))
	for i, result := range results {
		searchResults[i] = SearchResult{
			Rank:        i + 1,
			URL:         result.URL,
			Title:       result.Title,
			Description: result.Description,
		}
	}

	return searchResults, err
}

func (g Google) searchWithAPI(query string) ([]SearchResult, error) {
	if g.JsonKeyFilePath == "" || g.CustomCX == "" {
		return nil, errors.New("missing Service Account JSON key file path or Custom Search Engine ID")
	}

	ctx := context.Background()
	csService, err := customsearch.NewService(ctx, option.WithCredentialsFile(g.JsonKeyFilePath), option.WithScopes("https://www.googleapis.com/auth/cse"))
	if err != nil {
		return nil, fmt.Errorf("error creating custom search service: %v", err)
	}

	search := csService.Cse.List().Cx(g.CustomCX).Q(query)
	results, err := search.Do()
	if err != nil {
		return nil, fmt.Errorf("error executing search: %v", err)
	}

	searchResults := make([]SearchResult, len(results.Items))
	for i, result := range results.Items {
		searchResults[i] = SearchResult{
			Rank:        i + 1,
			URL:         result.Link,
			Title:       result.Title,
			Description: result.Snippet,
		}
	}

	return searchResults, nil
}

func (g Google) Search(query string) ([]SearchResult, error) {
	if g.JsonKeyFilePath == "" || g.CustomCX == "" {
		return g.searchAnonymous(query)
	}
	return g.searchWithAPI(query)
}
