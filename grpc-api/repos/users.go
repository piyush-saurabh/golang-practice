package repos

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/piyush-saurabh/golang-practice/grpc-api/types"
)

// UsersRepo - the user repo interface
// Add the functionalities here
type UsersRepo interface {
	Create(*types.User) error
	FindByID(int64) (*types.User, error)
	FindByEmail(string) (*types.User, error)
	Update(*types.User) error
}

type usersRepo struct {
	db *xorm.Engine
}

// NewUsersRepo - returns a new user repo
func NewUsersRepo(db *xorm.Engine) UsersRepo {
	return &usersRepo{db: db}
}

// Create - Add a new user to the database
func (u usersRepo) Create(user *types.User) (err error) {
	if err = types.Validate(user); err != nil {
		return
	}

	// Insert method comes from xorm
	if _, err = u.db.Insert(user); err != nil {
		return
	}
	return
}

// FindByID method accepts the ID and returns the User
func (u usersRepo) FindByID(id int64) (user *types.User, err error) {
	if id <= 0 {
		err := fmt.Errorf("Valid +ve ID is required to find a user")
		return nil, err
	}

	user = new(types.User)

	has, err := u.db.ID(id).Get(user)
	if err != nil {
		return
	}

	if !has {
		err = fmt.Errorf("Unable to find user")
		return
	}

	return

}

// FindByEmail method accepts an email ID and returns a User
func (u usersRepo) FindByEmail(email string) (user *types.User, err error) {
	if len(email) == 0 {
		err := fmt.Errorf("Valid email is required to find a user")
		return nil, err
	}

	user = new(types.User)
	user.Email = email

	has, err := u.db.Get(user)
	if err != nil {
		return
	}

	if !has {
		err = fmt.Errorf("Unable to find user")
		return
	}

	return
}

// Update - Update the user
func (u usersRepo) Update(user *types.User) (err error) {
	if user == nil || user.ID <= 0 {
		return fmt.Errorf("Invalid user passed")
	}

	if _, err := u.db.ID(user.ID).Update(user); err != nil {
		return err
	}
	return
}
