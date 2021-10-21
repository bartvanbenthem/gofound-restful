package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}
