package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/bartvanbenthem/gofound-restful/internal/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var app config.AppConfig

func TestMain(m *testing.M) {
	// change this to true when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := NewTestRepo(&app)
	NewHandlers(repo)

	os.Exit(m.Run())

}

func getRoutes() http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.StripSlashes)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(enableCORS)

	// public routes
	router.Group(func(router chi.Router) {
		router.Get("/status", Repo.Status)
		router.Get("/v1/software", Repo.GetAllSoftware)
		router.Get("/v1/software/{id}", Repo.GetSoftwareByID)
		router.Get("/v1/categories", Repo.GetAllCategories)
		router.Get("/v1/categories/{category_id}", Repo.GetAllSoftwareByCategory)

		router.Post("/v1/login", Repo.Login)
	})

	// protected routes
	router.Group(func(router chi.Router) {
		router.Use(TokenVerify)
		router.Get("/v1/admin/deletesoftware/{id}", Repo.DeleteSoftware)
		router.Post("/v1/admin/editSoftware", Repo.EditSoftware)
		router.Post("/v1/admin/signup", Repo.Signup)
	})

	return router
}

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
