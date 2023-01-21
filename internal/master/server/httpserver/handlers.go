package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/kozyrev-m/keeper/internal/master/model"
)

// handleRegisterUser creates new user in the system.
func (s *Server) handleRegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requestUser{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Login:    req.Login,
			Password: req.Password,
		}

		if err := s.store.CreateUser(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}
}