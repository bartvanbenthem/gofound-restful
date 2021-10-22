package dbrepo

import (
	"database/sql"

	"github.com/bartvanbenthem/gofound-restful/internal/config"
	"github.com/bartvanbenthem/gofound-restful/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
