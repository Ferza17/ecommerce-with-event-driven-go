package cassandradb

import (
	"github.com/gocql/gocql"

	"github.com/Ferza17/event-driven-product-service/module/product/repository/elasticsearch"
)

type productCassandraDBRepository struct {
	session *gocql.Session
}

func NewProductCassandraDBRepository(session *gocql.Session) elasticsearch.ProductElasticsearchRepositoryStore {
	return &productCassandraDBRepository{
		session: session,
	}
}
