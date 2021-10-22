package config

import (
	"log"

	"github.com/bartvanbenthem/gofound-restfull/utils"
)

// AppConfig holds the application config
type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Utils        *utils.JResponse
}
