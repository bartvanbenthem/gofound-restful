package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/bartvanbenthem/gofound-restful/internal/models"
)

func (m *postgresDBRepo) Signup(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into users (email, password,
				created_at, updated_at) values($1, $2, $3, $4)`

	_, err := m.DB.ExecContext(ctx, stmt,
		u.Email,
		u.Password,
		u.CreatedAt,
		u.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (m *postgresDBRepo) Login(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select * from users where email=$1"
	row := m.DB.QueryRowContext(ctx, query, email)

	var u models.User

	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}
	return u, err
}
