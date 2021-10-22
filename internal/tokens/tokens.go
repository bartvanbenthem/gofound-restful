package tokens

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bartvanbenthem/gofound-restful/internal/models"
	"github.com/bartvanbenthem/gofound-restful/internal/utils"
	"github.com/dgrijalva/jwt-go"
)

var j *utils.JResponse

func GenerateToken(user models.User) (string, error) {
	secret := "secret" //os.Getenv("GOFOUND_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "gofound.nl",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Printf("Error generating token: %v\n", err)
	}

	return tokenString, nil
}

func TokenVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte("secret"), nil
			})

			if err != nil {
				j.ErrorJSON(w, err)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				j.ErrorJSON(w, err)
				return
			}
		} else {
			err := fmt.Errorf("invalid token")
			j.ErrorJSON(w, err)
			return
		}
	})
}
