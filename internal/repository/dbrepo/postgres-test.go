package dbrepo

import (
	"github.com/bartvanbenthem/gofound-restful/internal/models"
)

// All returns all software and error, if any
func (m *testDBRepo) AllSoftware(category ...int) ([]*models.Software, error) {
	s := []*models.Software{}
	return s, nil
}

// Get returns one software program and error, if any
func (m *testDBRepo) GetSoftwareByID(id int) (*models.Software, error) {
	s := models.Software{}
	return &s, nil
}

func (m *testDBRepo) AllCategories() ([]*models.Category, error) {
	c := []*models.Category{}
	return c, nil
}

func (m *testDBRepo) InsertSoftware(s models.Software) error {
	return nil
}

func (m *testDBRepo) UpdateSoftware(s models.Software) error {
	return nil
}

func (m *testDBRepo) DeleteSoftware(id int) error {
	return nil
}

func (m *testDBRepo) Signup(u models.User) error {
	return nil
}

func (m *testDBRepo) Login(email string) (models.User, error) {
	u := models.User{}
	return u, nil
}
