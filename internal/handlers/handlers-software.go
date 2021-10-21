package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (m *Repository) GetAllSoftware(w http.ResponseWriter, r *http.Request) {
	software, err := m.DB.AllSoftware()
	if err != nil {
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	err = m.App.Utils.WriteJSON(w, http.StatusOK, software, "software")
	if err != nil {
		m.App.Utils.ErrorJSON(w, err)
		return
	}

}

func (m *Repository) GetSoftwareByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = errors.New("invalid id parameter")
		m.App.Utils.ErrorJSON(w, err)
		return
	}

	software, err := m.DB.GetSoftwareByID(id)

	err = m.App.Utils.WriteJSON(w, http.StatusOK, software, "software")
	if err != nil {
		m.App.Utils.ErrorJSON(w, err)
		return
	}
}
