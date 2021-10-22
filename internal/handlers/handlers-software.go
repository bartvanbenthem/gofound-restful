package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/bartvanbenthem/gofound-restful/internal/models"
	"github.com/go-chi/chi"
)

func (m *Repository) GetAllSoftware(w http.ResponseWriter, r *http.Request) {
	software, err := m.DB.AllSoftware()
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	err = m.App.Utils.WriteJSON(w, http.StatusOK, software, "software")
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

}

func (m *Repository) GetSoftwareByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = errors.New("invalid id parameter")
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	software, err := m.DB.GetSoftwareByID(id)

	err = m.App.Utils.WriteJSON(w, http.StatusOK, software, "software")
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}
}

func (m *Repository) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := m.DB.AllCategories()
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	err = m.App.Utils.WriteJSON(w, http.StatusOK, categories, "categories")
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}
}

func (m *Repository) GetAllSoftwareByCategory(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.Atoi(chi.URLParam(r, "category_id"))
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	software, err := m.DB.AllSoftware(categoryID)
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	err = m.App.Utils.WriteJSON(w, http.StatusOK, software, "software")
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}
}

func (m *Repository) DeleteSoftware(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	err = m.DB.DeleteSoftware(id)
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	ok := m.App.Utils.NewJResponse(true, "")

	err = m.App.Utils.WriteJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}
}

type SoftwarePayload struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"release_date"`
}

func (m *Repository) EditSoftware(w http.ResponseWriter, r *http.Request) {
	var payload SoftwarePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	var s models.Software

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := m.DB.GetSoftwareByID(id)
		s = *m
		s.UpdatedAt = time.Now()
	}

	s.ID, _ = strconv.Atoi(payload.ID)
	s.Name = payload.Name
	s.Description = payload.Description
	s.ReleaseDate, _ = time.Parse("2000-01-02", payload.ReleaseDate)
	s.Year = s.ReleaseDate.Year()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	if s.ID == 0 {
		err = m.DB.InsertSoftware(s)
		if err != nil {
			m.App.ErrorLog.Printf("%s\n", err)
			m.App.Utils.ErrorJSON(w, err)
			return
		}
	} else {
		err = m.DB.UpdateSoftware(s)
		if err != nil {
			m.App.ErrorLog.Printf("%s\n", err)
			m.App.Utils.ErrorJSON(w, err)
			return
		}
	}

	ok := m.App.Utils.NewJResponse(true, "")

	err = m.App.Utils.WriteJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		m.App.ErrorLog.Printf("%s\n", err)
		m.App.Utils.ErrorJSON(w, err)
		return
	}
}
