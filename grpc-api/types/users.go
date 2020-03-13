package types

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// xorm is used for database. xorm should know which is the primary key, autoincrement the value

// TempUser - Temporary user for creating a new user with confirm password field
type TempUser struct {
	ID              int64  `json:"id"`
	FirstName       string `json:"first_name" validate:"required,gte=4"`
	LastName        string `json:"last_name" validate:"required,gte=4"`
	Email           string `json:"email"`
	Password        string `json:"-" validate:"required,gte=8"`
	ConfirmPassword string `json:"-" validate:"required,gte=8"`
}

// User : This struct represents a user in the system
type User struct {
	ID        int64  `json:"id" xorm:"'id' pk autoincr" schema:"id"`
	FirstName string `json:"first_name" xorm:"first_name" schema:"first_name" validate:"required,gte=4"`
	LastName  string `json:"last_name" xorm:"last_name" schema:"last_name" validate:"required,gte=4"`
	Email     string `json:"email" xorm:"email" schema:"email" validate:"required,contains=@"`
	Password  string `json:"-" xorm:"password" schema:"password" validate:"required,gte=8"`
	Visible   bool   `json:"visible" xorm:"visible" schema:"visible"`
}

// TableName method returns the table when using xorm
func (u *User) TableName() string {
	return "users"
}

// NewUser method creates a new user from temproary user if both passwords matches
func NewUser(newUser *TempUser) (user *User, err error) {

	if newUser.Password != newUser.ConfirmPassword {
		err := fmt.Errorf("Password and Confirm Password do not match")
		return nil, err
	}

	// Initializing the struct
	user = &User{
		// ID is auto-generated
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Visible:   true,
	}

	if err = user.SetPassword(newUser.Password); err != nil {
		return nil, err
	}
	return user, nil
}

// SetPassword method uses bcrypt to hash the password
func (u *User) SetPassword(password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

// Authenticate method authenticates the password against the stored hash
func (u *User) Authenticate(password string) error {
	if !u.Visible {
		return fmt.Errorf("User is inactive")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}
