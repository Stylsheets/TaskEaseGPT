package pinecone

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/pinecone-io/go-pinecone/pinecone_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	conn   *grpc.ClientConn
	client pinecone_grpc.VectorServiceClient
	ctx    context.Context
}

func Init(indexName, projectName, pineconeEnv, apiKey string) (*Client, error) {
	config := &tls.Config{}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "api-key", apiKey)
	target := fmt.Sprintf("%s-%s.svc.%s.pinecone.io:443", indexName, projectName, pineconeEnv)
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(credentials.NewTLS(config)),
		grpc.WithAuthority(target),
	)
	if err != nil {
		return nil, err
	}

	client := pinecone_grpc.NewVectorServiceClient(conn)

	return &Client{conn: conn, client: client, ctx: ctx}, nil
}

func (pc *Client) Close() error {
	err := pc.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (pc *Client) Upsert(vectors []*pinecone_grpc.Vector, namespace string) (*pinecone_grpc.UpsertResponse, error) {
	return pc.client.Upsert(pc.ctx, &pinecone_grpc.UpsertRequest{
		Vectors:   vectors,
		Namespace: namespace,
	})
}

func (pc *Client) Fetch(ids []string, namespace string) (*pinecone_grpc.FetchResponse, error) {
	return pc.client.Fetch(pc.ctx, &pinecone_grpc.FetchRequest{
		Ids:       ids,
		Namespace: namespace,
	})
}

func (pc *Client) Query(queries []*pinecone_grpc.QueryVector, topK uint32, includeValues bool, namespace string) (*pinecone_grpc.QueryResponse, error) {
	return pc.client.Query(pc.ctx, &pinecone_grpc.QueryRequest{
		Queries:       queries,
		TopK:          topK,
		IncludeValues: includeValues,
		Namespace:     namespace,
	})
}

func (pc *Client) Delete(ids []string, namespace string) (*pinecone_grpc.DeleteResponse, error) {
	return pc.client.Delete(pc.ctx, &pinecone_grpc.DeleteRequest{
		Ids:       ids,
		Namespace: namespace,
	})
}

func (pc *Client) DescribeIndexStats() (*pinecone_grpc.DescribeIndexStatsResponse, error) {
	return pc.client.DescribeIndexStats(pc.ctx, &pinecone_grpc.DescribeIndexStatsRequest{})
}
