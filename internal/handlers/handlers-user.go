package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/bartvanbenthem/gofound-restful/internal/config"
	"github.com/bartvanbenthem/gofound-restful/internal/models"
	"github.com/bartvanbenthem/gofound-restful/internal/tokens"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var jwt config.JWT
	var err error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		err := errors.New("email is missing")
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	if user.Password == "" {
		err := errors.New("password is missing")
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	password := user.Password

	user, err = m.DB.Login(user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err := errors.New("sql login error")
			m.App.Utils.ErrorJSON(w, err)
			return
		} else {
			log.Fatal(err)
		}
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		err := errors.New("error comparing hash and password")
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	token, err := tokens.GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	m.App.Utils.WritePlainJSON(w, jwt)
}

func (m *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		err := errors.New("email is missing")
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	if user.Password == "" {
		err := errors.New("password is missing")
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hash)

	err = m.DB.Signup(user)

	if err != nil {
		err := errors.New("server error")
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	user.Password = ""

	err = m.App.Utils.WriteJSON(w, http.StatusOK, user, "response")
	if err != nil {
		m.App.Utils.ErrorJSON(w, err)
		return
	}
}
