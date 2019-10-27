package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)
var mySigningKey = []byte("captainjacksparrowsayshi")

func Use(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware")
		c, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(c.Value) != 0 {
			token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				_, _ = fmt.Fprintf(w, err.Error())
				log.Println(err)
			}

			if token.Valid {
				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r)
			}
		} else {
			_, err := fmt.Fprintf(w, "Not Authorized")
			log.Println(err)
		}
	})
}
// Currently the payload is just the username
func GenerateJWT(payload string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = payload
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatal("Something Went Wrong:", err.Error())
		return "", err
	}

	return tokenString, nil
}
/*
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			_, _ = fmt.Fprintf(w, err.Error())
			log.Println(err)
		}

		if token.Valid {
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		}
	}
 */