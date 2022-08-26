package search

import (
	"bytes"
	"context"
	"encoding/json"
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/solrac97gr/feed-cqrs/models"
)

type ElasticSearchRepository struct {
	client *elastic.Client
}

func NewElasticSearchRepository(url string) (*ElasticSearchRepository, error) {
	client, err := elastic.NewClient(
		elastic.Config{
			Addresses: []string{url},
		},
	)
	if err != nil {
		return nil, err
	}
	return &ElasticSearchRepository{client: client}, nil
}

var _ Repository = (*ElasticSearchRepository)(nil)

func (e ElasticSearchRepository) Close() {
	//
}

func (e ElasticSearchRepository) IndexFeed(ctx context.Context, feed *models.Feed) error {
	body, _ := json.Marshal(feed)
	_, err := e.client.Index(
		"feeds",
		bytes.NewReader(body),
		e.client.Index.WithDocumentID(feed.ID),
		e.client.Index.WithContext(ctx),
		e.client.Index.WithRefresh("wait_for"),
	)
	return err
}

func (e ElasticSearchRepository) SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	//TODO implement me
	panic("implement me")
}
