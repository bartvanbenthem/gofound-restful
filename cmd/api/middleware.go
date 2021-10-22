package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next.ServeHTTP(w, r)
	})
}

func TokenVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					app.ErrorLog.Printf("parsing token")
					return nil, fmt.Errorf("There was an error")
				}

				return []byte("secret"), nil
			})

			if err != nil {
				app.ErrorLog.Printf("%s\n", err)
				app.Utils.ErrorJSON(w, err)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				app.InfoLog.Printf("%s\n", err)
				app.Utils.ErrorJSON(w, err)
				return
			}
		} else {
			err := fmt.Errorf("invalid token")
			app.InfoLog.Printf("%s\n", err)
			app.Utils.ErrorJSON(w, err)
			return
		}
	})
}
