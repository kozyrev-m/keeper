package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/kozyrev-m/keeper/internal/master/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"errors"
)

func TestServer_HandleRegisterUser(t *testing.T) {
	store := &mockStore{}

	s := New(store, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		initMock     func()
		expectedCode int
	} {
		{
			name: "valid",
			payload: map[string]string{
				"login":    "user",
				"password": "password",
			},
			initMock: func() {
				store.createUser = func(m *model.User) error {
					m.ID = 1
					return nil
				}
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			initMock: func() {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "can't insert user",
			payload: map[string]string{
				"login":    "user",
				"password": "password",
			},
			initMock: func() {
				store.createUser = func(m *model.User) error {
					return errors.New("can't insert user -- something bad happened")
				}
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.initMock()

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}

			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleCreateSession(t *testing.T) {
	u := model.TestUser(t)

	store := &mockStore{}

	s := New(store, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name    string
		payload interface{}
		initMock func()
		expectedCode int
	} {
		{
			name: "valid",
			payload: map[string]string {
				"login": u.Login,
				"password": u.Password,
			},
			initMock: func() {
				store.findUserByLogin = func(login string) (*model.User, error) {
					user := &model.User {
						Login: u.Login,
						Password: u.Password,
					}
					
					user.BeforeCreate()
					user.ID = 1

					require.Equal(t, login, user.Login)

					return user, nil
				}
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			initMock: func() {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "incorrect login",
			payload: map[string]string{
				"login": "incorrect",
				"password": u.Password,
			},
			initMock: func() {
				store.findUserByLogin = func(login string) (*model.User, error) {
					return nil, errors.New("not find user")
				}
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "incorrect password",
			payload: map[string]string{
				"login": u.Login,
				"password": "incorrect",
			},
			initMock: func() {
				store.findUserByLogin = func(login string) (*model.User, error) {
					user := &model.User {
						Login: u.Login,
						Password: u.Password,
					}
					
					user.BeforeCreate()
					user.ID = 1

					return user, nil
				}
			},
			expectedCode: http.StatusUnauthorized,
		},
	}


	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.initMock()

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}