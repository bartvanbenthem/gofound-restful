package dbrepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bartvanbenthem/gofound-restful/internal/models"
)

// All returns all software and error, if any
func (m *postgresDBRepo) AllSoftware(category ...int) ([]*models.Software, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(category) > 0 {
		where = fmt.Sprintf("where id in (select software_id from software_categories where category_id = %d)", category[0])
	}

	query := fmt.Sprintf(`select id, name, description, year, release_date,
				created_at, updated_at from software  %s order by name`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slist []*models.Software

	for rows.Next() {
		var s models.Software
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Description,
			&s.Year,
			&s.ReleaseDate,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// get categories, if any
		categorieQuery := `select
			sc.id, sc.software_id, sc.category_id, c.category_name
		from
			software_categories sc
			left join categories c on (c.id = sc.category_id)
		where
			sc.software_id = $1
		`

		categoryRows, _ := m.DB.QueryContext(ctx, categorieQuery, s.ID)

		categories := make(map[int]string)
		for categoryRows.Next() {
			var sc models.SoftwareCategory
			err := categoryRows.Scan(
				&sc.ID,
				&sc.SoftwareID,
				&sc.CategoryID,
				&sc.Category.CategoryName,
			)
			if err != nil {
				return nil, err
			}
			categories[sc.ID] = sc.Category.CategoryName
		}
		categoryRows.Close()

		s.SoftwareCategory = categories
		slist = append(slist, &s)

	}
	return slist, nil
}

// Get returns one software program and error, if any
func (m *postgresDBRepo) GetSoftwareByID(id int) (*models.Software, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, name, description, year, release_date,
				created_at, updated_at from software where id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	var s models.Software

	err := row.Scan(
		&s.ID,
		&s.Name,
		&s.Description,
		&s.Year,
		&s.ReleaseDate,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// get categories, if any
	query = `select
				sc.id, sc.software_id, sc.category_id, c.category_name
			from
				software_categories sc
				left join categories c on (c.id = sc.category_id)
			where
				sc.software_id = $1
	`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	categories := make(map[int]string)
	for rows.Next() {
		var sc models.SoftwareCategory
		err := rows.Scan(
			&sc.ID,
			&sc.CategoryID,
			&sc.SoftwareID,
			&sc.Category.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		categories[sc.ID] = sc.Category.CategoryName
	}

	s.SoftwareCategory = categories

	return &s, nil
}

func (m *postgresDBRepo) AllCategories() ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, category_name, created_at, updated_at from categories order by category_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clist []*models.Category

	for rows.Next() {
		var c models.Category
		err := rows.Scan(
			&c.ID,
			&c.CategoryName,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		clist = append(clist, &c)
	}

	return clist, nil
}

func (m *postgresDBRepo) InsertSoftware(s models.Software) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into software (name, description, year, release_date,
				created_at, updated_at) values ($1, $2, $3, $4, $5, $6)`

	_, err := m.DB.ExecContext(ctx, stmt,
		s.Name,
		s.Description,
		s.Year,
		s.ReleaseDate,
		s.CreatedAt,
		s.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *postgresDBRepo) UpdateSoftware(s models.Software) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update software set name = $1, description = $2, year = $3, release_date = $4, 
				updated_at = $5 where id = $6`

	_, err := m.DB.ExecContext(ctx, stmt,
		s.Name,
		s.Description,
		s.Year,
		s.ReleaseDate,
		s.UpdatedAt,
		s.ID,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *postgresDBRepo) DeleteSoftware(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "delete from software where id = $1"

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
