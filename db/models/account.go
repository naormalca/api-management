package models

import (
	jwt "github.com/naormalca/api-management/api/middleware"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Account struct {
	gorm.Model

	Username string `gorm:"type:varchar(100); unique; not null"`
	Password string `gorm:"not null"`
	//Role     Role   `gorm:"type:varchar(5); not null"`
	//Active   bool   `gorm:"not null"`
	Token    string `gorm:"not null"`
}
func (account *Account) PrepareAccount() error {
	password, err := HashPassword(account.Password)
	if err != nil {
		return err
	}
	account.Password = password
	account.Token, err = jwt.GenerateJWT(account.Username)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

