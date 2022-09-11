package repository

import "github.com/gocql/gocql"

type userCassandraDBRepository struct {
	session *gocql.Session
}

func NewUserCassandraDBRepository(session *gocql.Session) UserCassandraDBRepositoryStore {
	return &userCassandraDBRepository{
		session: session,
	}
}
