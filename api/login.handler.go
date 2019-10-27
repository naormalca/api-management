package api

import (
	"github.com/naormalca/api-management/db"
	"github.com/naormalca/api-management/db/models"
	"encoding/json"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	accountReq := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(accountReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Unmarshell the body failed!")
		return
	}
	// Fetch the account details from DB
	account, err := db.DBService.Account.Find(accountReq.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("The user not exists!")
		return
	}
	err = models.CheckPasswordHash(accountReq.Password, account.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Bad password")
		log.Println(err)
		return
	}
	log.Println(account.Username, "Logged")
	return
}
