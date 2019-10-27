package api

import (
	"github.com/naormalca/api-management/db"
	"github.com/naormalca/api-management/db/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Unmarshell the body failed!")
		return
	}
	fmt.Println(account)
	err = account.PrepareAccount()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.DBService.Account.Find(account.Username)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("The user name exists!")
		return
	}
	err = db.DBService.Account.Create(account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:       "jwt",
		Value:      account.Token,
	})
	log.Println("Account created!")
}