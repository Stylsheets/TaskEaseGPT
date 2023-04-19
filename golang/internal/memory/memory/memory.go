package memory

import (
	"TaskEaseGPT/internal/memory/pinecone"
)

type MemoryAdapter interface {
	DescribeIndexStats(options *DescribeIndexStatsOptions) (*DescribeIndexStatsOutput, error)
	Query(options *QueryOptions) (*QueryOutput, error)
	Delete(options *DeleteOptions) (*DeleteOutput, error)
	Fetch(options *FetchOptions) (*FetchOutput, error)
	Update(options *UpdateOptions) (*UpdateOutput, error)
	Upsert(options *UpsertOptions) (*UpsertOutput, error)
}

type PineconeAdapter struct {
	Client *pinecone.Client
}

func NewPineconeAdapter(apiKey, baseURL string) *PineconeAdapter {
	return &PineconeAdapter{
		Client: pinecone.NewClient(apiKey, baseURL),
	}
}

// DescribeIndexStats

type DescribeIndexStatsOptions struct {
	Filter map[string]string
}

type DescribeIndexStatsOutput struct {
	Namespaces       map[string]struct{ VectorCount int64 }
	Dimension        int64
	IndexFullness    float32
	TotalVectorCount int64
}

func (pa *PineconeAdapter) DescribeIndexStats(options *DescribeIndexStatsOptions) (*DescribeIndexStatsOutput, error) {
	res, err := pa.Client.DescribeIndexStats(&pinecone.DescribeIndexStatsRequest{
		Filter: options.Filter,
	})

	if err != nil {
		return nil, err
	}

	return &DescribeIndexStatsOutput{
		Namespaces:       res.Namespaces,
		Dimension:        res.Dimension,
		IndexFullness:    res.IndexFullness,
		TotalVectorCount: res.TotalVectorCount,
	}, nil
}

// Query

type QueryOptions struct {
	Namespace       string
	TopK            int64
	Filter          map[string]string
	IncludeValues   bool
	IncludeMetadata bool
	Vector          []float32
	SparseVector    *pinecone.SparseVector
	ID              string
}

type QueryOutput struct {
	Matches   []pinecone.QueryMatch
	Namespace string
}

func (pa *PineconeAdapter) Query(options *QueryOptions) (*QueryOutput, error) {
	res, err := pa.Client.Query(&pinecone.QueryRequest{
		Namespace:       options.Namespace,
		TopK:            options.TopK,
		Filter:          options.Filter,
		IncludeValues:   options.IncludeValues,
		IncludeMetadata: options.IncludeMetadata,
		Vector:          options.Vector,
		SparseVector:    options.SparseVector,
		ID:              options.ID,
	})

	if err != nil {
		return nil, err
	}

	return &QueryOutput{
		Matches:   res.Matches,
		Namespace: res.Namespace,
	}, nil
}

// Delete

type DeleteOptions struct {
	IDs       []string
	DeleteAll bool
	Namespace string
	Filter    map[string]string
}

type DeleteOutput struct{}

func (pa *PineconeAdapter) Delete(options *DeleteOptions) (*DeleteOutput, error) {
	_, err := pa.Client.Delete(&pinecone.DeleteRequest{
		IDs:       options.IDs,
		DeleteAll: options.DeleteAll,
		Namespace: options.Namespace,
		Filter:    options.Filter,
	})

	if err != nil {
		return nil, err
	}

	return &DeleteOutput{}, nil
}

// Fetch

type FetchOptions struct {
	IDs       []string
	Namespace string
}

type FetchOutput struct {
	Vectors   map[string]pinecone.FetchVector
	Namespace string
}

func (pa *PineconeAdapter) Fetch(options *FetchOptions) (*FetchOutput, error) {
	res, err := pa.Client.Fetch(&pinecone.FetchRequest{
		IDs:       options.IDs,
		Namespace: options.Namespace,
	})

	if err != nil {
		return nil, err
	}

	return &FetchOutput{
		Vectors:   res.Vectors,
		Namespace: res.Namespace,
	}, nil
}

// Update

type UpdateOptions struct {
	ID           string
	Values       []float32
	SparseValues *pinecone.SparseVector
	SetMetadata  map[string]string
	Namespace    string
}

type UpdateOutput struct{}

func (pa *PineconeAdapter) Update(options *UpdateOptions) (*UpdateOutput, error) {
	_, err := pa.Client.Update(&pinecone.UpdateRequest{
		ID:           options.ID,
		Values:       options.Values,
		SparseValues: options.SparseValues,
		SetMetadata:  options.SetMetadata,
		Namespace:    options.Namespace,
	})

	if err != nil {
		return nil, err
	}

	return &UpdateOutput{}, nil
}

// Upsert

type UpsertOptions struct {
	Vectors   []pinecone.UpsertVector
	Namespace string
}

type UpsertOutput struct {
	UpsertedCount int64
}

func (pa *PineconeAdapter) Upsert(options *UpsertOptions) (*UpsertOutput, error) {
	res, err := pa.Client.Upsert(&pinecone.UpsertRequest{
		Vectors:   options.Vectors,
		Namespace: options.Namespace,
	})

	if err != nil {
		return nil, err
	}

	return &UpsertOutput{
		UpsertedCount: res.UpsertedCount,
	}, nil
}
