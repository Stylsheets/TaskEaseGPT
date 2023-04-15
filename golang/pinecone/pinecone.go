package pinecone

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

type ErrorResponse struct {
	Code    int32
	Message string
	Details []struct {
		TypeUrl string
		Value   string
	}
}

type HTTPError struct {
	StatusCode int
	ErrorResponse
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s, Details: %v", e.StatusCode, e.Message, e.Details)
}

func NewClient(apiKey, baseURL string) *Client {
	return &Client{
		client:  &http.Client{},
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (c *Client) doRequest(req *http.Request, result interface{}) error {
	req.Header.Set("Api-Key", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return err
		}
		return &HTTPError{
			StatusCode:    resp.StatusCode,
			ErrorResponse: errorResponse,
		}
	}

	if result != nil {
		err = json.NewDecoder(resp.Body).Decode(result)
	}
	return err
}

// DescribeIndexStats

type DescribeIndexStatsRequest struct {
	Filter map[string]string `json:"filter,omitempty"`
}

type DescribeIndexStatsResponse struct {
	Namespaces       map[string]struct{ VectorCount int64 } `json:"namespaces"`
	Dimension        int64                                  `json:"dimension"`
	IndexFullness    float64                                `json:"index_fullness"`
	TotalVectorCount int64                                  `json:"totalVectorCount"`
}

func (c *Client) DescribeIndexStats(req *DescribeIndexStatsRequest) (*DescribeIndexStatsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/describe_index_stats", c.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")

	var resp DescribeIndexStatsResponse
	err = c.doRequest(httpReq, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Query

type QueryRequest struct {
	Namespace       string            `json:"namespace"`
	TopK            int64             `json:"topK"`
	Filter          map[string]string `json:"filter,omitempty"`
	IncludeValues   bool              `json:"includeValues"`
	IncludeMetadata bool              `json:"includeMetadata"`
	Vector          []float64         `json:"vector"`
	SparseVector    *SparseVector     `json:"sparseVector,omitempty"`
	ID              string            `json:"id"`
}

type SparseVector struct {
	Indices []int64   `json:"indices"`
	Values  []float64 `json:"values"`
}

type QueryMatch struct {
	ID           string                 `json:"id"`
	Score        float64                `json:"score"`
	Values       []float64              `json:"values,omitempty"`
	SparseValues *SparseVector          `json:"sparseValues,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type QueryResponse struct {
	Matches   []QueryMatch `json:"matches"`
	Namespace string       `json:"namespace"`
}

func (c *Client) Query(req *QueryRequest) (*QueryResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/query", c.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")

	var resp QueryResponse
	err = c.doRequest(httpReq, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Delete

type DeleteRequest struct {
	IDs       []string          `json:"ids,omitempty"`
	DeleteAll bool              `json:"deleteAll"`
	Namespace string            `json:"namespace"`
	Filter    map[string]string `json:"filter,omitempty"`
}

type DeleteResponse struct{}

func (c *Client) Delete(req *DeleteRequest) (*DeleteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/delete", c.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")

	var resp DeleteResponse
	err = c.doRequest(httpReq, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Fetch

type FetchRequest struct {
	IDs       []string
	Namespace string
}

type FetchVector struct {
	ID           string                 `json:"id"`
	Values       []float64              `json:"values,omitempty"`
	SparseValues *SparseVector          `json:"sparseValues,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type FetchResponse struct {
	Vectors   map[string]FetchVector `json:"vectors"`
	Namespace string                 `json:"namespace"`
}

func (c *Client) Fetch(req *FetchRequest) (*FetchResponse, error) {
	query := url.Values{}
	for _, id := range req.IDs {
		query.Add("ids", id)
	}

	httpReq, err := http.NewRequest("GET", fmt.Sprintf("%s/vectors/fetch?%s", c.baseURL, query.Encode()), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")

	var resp FetchResponse
	err = c.doRequest(httpReq, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Update

type UpdateRequest struct {
	ID           string            `json:"id"`
	Values       []float64         `json:"values,omitempty"`
	SparseValues *SparseVector     `json:"sparseValues,omitempty"`
	SetMetadata  map[string]string `json:"setMetadata,omitempty"`
	Namespace    string            `json:"namespace"`
}

type UpdateResponse struct{}

func (c *Client) Update(req *UpdateRequest) (*UpdateResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/update", c.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")

	var resp UpdateResponse
	err = c.doRequest(httpReq, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Upsert

type UpsertRequest struct {
	Vectors   []UpsertVector `json:"vectors"`
	Namespace string         `json:"namespace"`
}

type UpsertVector struct {
	ID           string            `json:"id"`
	Values       []float64         `json:"values"`
	SparseValues *SparseVector     `json:"sparseValues,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type UpsertResponse struct {
	UpsertedCount int64 `json:"upsertedCount"`
}

func (c *Client) Upsert(req *UpsertRequest) (*UpsertResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/vectors/upsert", c.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")

	var resp UpsertResponse
	err = c.doRequest(httpReq, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
