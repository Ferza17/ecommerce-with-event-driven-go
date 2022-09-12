package cassandradb

import "github.com/gocql/gocql"

type cartCassandraDBRepository struct {
	session *gocql.Session
}

func NewCartCassandraDBRepository(session *gocql.Session) CartCassandraDBRepositoryStore {
	return &cartCassandraDBRepository{
		session: session,
	}
}
