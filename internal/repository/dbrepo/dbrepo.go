package dbrepo

import (
	"database/sql"

	"github.com/bartvanbenthem/gofound-restful/internal/config"
	"github.com/bartvanbenthem/gofound-restful/internal/repository"
)

//PostgreSQL Database Repository
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

//Test Database Repository
type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
