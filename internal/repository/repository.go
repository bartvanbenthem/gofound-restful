package repository

import "github.com/bartvanbenthem/gofound-restfull/internal/models"

type DatabaseRepo interface {
	AllSoftware(category ...int) ([]*models.Software, error)
	GetSoftwareByID(id int) (*models.Software, error)
	AllCategories() ([]*models.Category, error)
	InsertSoftware(s models.Software) error
	UpdateSoftware(s models.Software) error
	DeleteSoftware(id int) error
	Login(email string) (models.User, error)
	Signup(u models.User) error
}
