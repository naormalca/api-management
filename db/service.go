package db

import (
	"github.com/naormalca/api-management/db/handlers"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Account  *handlers.AccountHandler
}

func NewService(db *gorm.DB) Service {
	return Service{
		Account:  handlers.NewAccountHandler(db),
	}
}