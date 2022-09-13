package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type productElasticsearchRepository struct {
	client *elasticsearch.Client
}

func NewProductElasticsearchRepository(client *elasticsearch.Client) ProductElasticsearchRepositoryStore {
	return &productElasticsearchRepository{
		client: client,
	}
}
