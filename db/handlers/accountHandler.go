package handlers

import "github.com/jinzhu/gorm"
import "github.com/naormalca/api-management/db/models"

type AccountHandler struct {
	db *gorm.DB
}

func NewAccountHandler(db *gorm.DB) *AccountHandler {
	return &AccountHandler{
		db,
	}
}

func (h *AccountHandler) Find(username string) (*models.Account, error) {
	var res models.Account

	if err := h.db.Find(&res, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (h *AccountHandler) FindBy(cond *models.Account) (*models.Account, error) {
	var res models.Account

	if err := h.db.Find(&res, cond).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (h *AccountHandler) Update(account *models.Account, accountID uint) error {
	return h.db.Model(models.Account{}).Where(" id = ? ", accountID).Update(account).Error
}

func (h *AccountHandler) Create(account *models.Account) error {
	return h.db.Create(account).Error
}
