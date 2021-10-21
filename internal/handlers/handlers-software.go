package handlers

import (
	"net/http"
)

func (m *Repository) GetAllSoftware(w http.ResponseWriter, r *http.Request) {
	software, err := m.DB.AllSoftware()
	if err != nil {
		m.errorJSON(w, err)
		return
	}

	err = m.writeJSON(w, http.StatusOK, software, "software")
	if err != nil {
		m.errorJSON(w, err)
		return
	}

}
