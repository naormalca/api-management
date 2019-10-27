package main

import (
	"github.com/naormalca/api-management/api"
	"github.com/naormalca/api-management/api/middleware"
	"github.com/naormalca/api-management/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.Load()
	r := mux.NewRouter()
	mw := []func(http.Handler) http.Handler{middleware.AuthRequired}
	r.HandleFunc("/register", api.RegisterHandler).
		Methods("POST")
	r.Handle("/login", middleware.Use(http.HandlerFunc(api.LoginHandler), mw...)).
		Methods("POST")

	defer db.Close()
	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":8081", r))

}


