package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/kozyrev-m/keeper/internal/master/model"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticateUser(t *testing.T) {
	sessionName := "session"

	u := model.TestUser(t)
	u.ID = 1

	store := &mockStore{}

	testCases := []struct {
		name string
		cookieValue map[interface{}]interface{}
		initMock func ()
		expectedCode int
	} {
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			initMock: func() {
				store.findUserByID = func(id int) (*model.User, error) {
					return u, nil
				}
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "not authenticated",
			cookieValue: nil,
			initMock: func ()  {},
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	s := New(store, sessions.NewCookieStore(secretKey))

	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		tc.initMock()

		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)

			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}