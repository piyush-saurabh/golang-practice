package repos_test

import (
	"database/sql"
	"testing"

	"github.com/go-xorm/xorm"
)

var (
	err   error
	db    *xorm.Engine
	dbSql *sql.DB
	mock  sqlmock.Sqlmock

	gr rep
)

func TestRepos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repos Suite")
}
