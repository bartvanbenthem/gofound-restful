package config

import (
	"log"

	"github.com/bartvanbenthem/gofound-restful/internal/utils"
)

// JWT holds the auth token
type JWT struct {
	Token string `json:"token"`
}

// AppConfig holds the application config
type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Utils        *utils.JResponse
	JWT          struct {
		Token string `json:"token"`
	}
}
