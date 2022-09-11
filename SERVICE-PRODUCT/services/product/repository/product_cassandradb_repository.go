package repository

import "github.com/gocql/gocql"

type productCassandraDBRepository struct {
	session *gocql.Session
}

func NewProductCassandraDBRepository(session *gocql.Session) ProductElasticsearchRepositoryStore {
	return &productCassandraDBRepository{
		session: session,
	}
}
