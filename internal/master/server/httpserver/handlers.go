package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/kozyrev-m/keeper/internal/master/model/usermodel"
	"github.com/kozyrev-m/keeper/internal/master/storage/store/filestorage"
)

// handleRegisterUser creates new user in the system.
func (s *Server) handleRegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requestUser{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &usermodel.User{
			Login:    req.Login,
			Password: req.Password,
		}

		if err := s.store.CreateUser(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()

		s.respond(w, r, http.StatusCreated, u)
	}
}

// handleCreateSession creates user session.
func (s *Server) handleCreateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requestUser{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.FindUserByLogin(req.Login)

		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectLoginOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID

		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, fmt.Sprintf("session successfully created for the user '%s'", u.Login))
	}
}

// handeWhoami gets information about user.
func (s *Server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*usermodel.User))
	}
}

// handleAddText adds some text.
func (s *Server) handleCreateText() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requestText{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		content := &datamodel.Text{
			BasePart: datamodel.BasePart{
				OwnerID:  u.ID,
				TypeID:   1,
				Metadata: req.Metadata,
			},
			Value: req.Text,
		}

		if err := s.store.CreateDataRecord(content); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, req)
	}
}

// handleGetTexts gets all user texts.
func (s *Server) handleGetTexts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		texts, err := s.store.FindTextsByOwner(u.ID)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		b, err := json.Marshal(texts)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, string(b))
	}
}

// handleAddPair adds login-password pair.
func (s *Server) handleAddPair() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requestPair{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		content := &datamodel.LoginPassword{
			BasePart: datamodel.BasePart{
				OwnerID:  u.ID,
				TypeID:   datamodel.TypePair,
				Metadata: req.Metadata,
			},
			Login:    req.Login,
			Password: req.Password,
		}

		if err := s.store.CreateDataRecord(content); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, req)
	}
}

// handleGetPairs gets login-password pairs.
func (s *Server) handleGetPairs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		pairs, err := s.store.FindPairsByOwner(u.ID)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		b, err := json.Marshal(pairs)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, string(b))
	}
}

// handleAddBankCard adds data of bank card.
func (s *Server) handleAddBankCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requestBankCard{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		content := &datamodel.BankCard {
			BasePart: datamodel.BasePart{
				OwnerID:  u.ID,
				TypeID:   datamodel.TypeBank,
				Metadata: req.Metadata,
			},
			PAN: req.PAN,
			CVV: req.CVV,
			ValidThru: req.ValidThru,
			Name: req.Name,
		}

		if err := s.store.CreateDataRecord(content); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, req)
	}
}

// handleGetBankCards gets bank card data.
func (s *Server) handleGetBankCards() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		cards, err := s.store.FindBankCardsByOwner(u.ID)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		b, err := json.Marshal(cards)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, string(b))
	}
}

// handleSaveFile saves user file.
func (s *Server) handleSaveFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		file, fheader, err := r.FormFile("file")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := s.store.CreateFile(u.ID, r.FormValue("metadata"), fheader.Filename, file); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, "")
	}
}

// handleFileList gets file list.
func (s *Server) handleFileList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		fileList, err := s.store.GetFileList(u.ID)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		b, err := json.Marshal(fileList)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, string(b))
	}
}

// handleDownloadFile get user file.
func (s *Server) handleDownloadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*usermodel.User)

		vars := mux.Vars(r)
		filename := vars["filename"]

		filepath := fmt.Sprintf("%s/%d/%s", filestorage.Dir, u.ID, filename)

		if !filestorage.ExistFile(filepath) {
			http.Error(w, errFileNotExist.Error(), http.StatusNoContent)
			return
		}

		http.ServeFile(w, r, filepath)
	}
}
