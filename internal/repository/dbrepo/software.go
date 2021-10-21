package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/bartvanbenthem/gofound-restfull/internal/models"
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
