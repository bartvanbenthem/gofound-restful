package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bartvanbenthem/gofound-restful/internal/models"
	"github.com/dgrijalva/jwt-go"
)

type JResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (j *JResponse) NewJResponse(ok bool, message string) *JResponse {
	r := JResponse{OK: ok, Message: message}
	return &r
}

func (j *JResponse) WritePlainJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func (j *JResponse) WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (j *JResponse) ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	j.WriteJSON(w, statusCode, theError, "error")
}

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
