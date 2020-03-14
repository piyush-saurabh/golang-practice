package repos

import (
	"github.com/go-xorm/xorm"
)

var globalRepositoryInstance *globalRepository

// GlobalRepository interface is used for testing
// Returns any build in repo/model
type GlobalRepository interface {
	Users() UsersRepo
}

// GlobalRepo - It is the global repository function
func GlobalRepo(db *xorm.Engine) GlobalRepository {
	if globalRepositoryInstance != nil {
		return globalRepositoryInstance
	}

	globalRepositoryInstance = &globalRepository{
		db:    db,
		repos: make(map[string]interface{}),
	}
	return globalRepositoryInstance
}

// This structure acts as the interface for the database
type globalRepository struct {
	db    *xorm.Engine
	repos map[string]interface{}
}

func (r *globalRepository) repoFactory(key string, factory func() interface{}) interface{} {
	if iface, exists := r.repos[key]; exists {
		return iface
	}
	iface := factory()
	r.repos[key] = iface
	return iface
}

func (r globalRepository) Users() UsersRepo {
	return r.repoFactory("Users", func() interface{} { return NewUsersRepo(r.db) }).(UsersRepo)
}
