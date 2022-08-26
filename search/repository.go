package search

import (
	"context"
	"github.com/solrac97gr/feed-cqrs/models"
)

type Repository interface {
	Close()
	IndexFeed(ctx context.Context, feed *models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]models.Feed, error)
}

var repo Repository

func SetSearchRepository(r Repository) {
	repo = r
}

func Close() {
	repo.Close()
}

func IndexFeed(ctx context.Context, feed *models.Feed) error {
	return repo.IndexFeed(ctx, feed)
}

func SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	return repo.SearchFeed(ctx, query)
}
