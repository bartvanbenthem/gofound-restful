package repository

import "github.com/bartvanbenthem/gofound-restfull/internal/models"

type DatabaseRepo interface {
	AllSoftware(category ...int) ([]*models.Software, error)
}
