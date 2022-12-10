package model

import (
	"github.com/ybkuroki/go-webapp-sample/repository"
	"golang.org/x/crypto/bcrypt"
)

// User defines struct of user data.
type User struct {
	user_id   uint   `gorm:"primary_key" json:"id"`
	user_name string `json:"user_name"`
	password  string `json:"-"`
}

const selectUser = "select u.user_id as id, u.user_name as name, .password as password from users u"

// TableName returns the table name of User struct and it is used by gorm.
func (User) TableName() string {
	return "users"
}

// NewUser is constructor.
func NewUser(user_name string, password string) *User {
	return &User{user_name: user_name, password: password}
}

// NewUserWithPlainPassword is constructor. And it is encoded plain text password by using bcrypt.
func NewUserWithPlainPassword(user_name string, password string, authorityID uint) *User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return &User{user_name: user_name, password: string(hashed)}
}

// Create persists this User data.
func (u *User) Create(rep repository.Repository) (*User, error) {
	if err := rep.Select("user_name", "password").Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
