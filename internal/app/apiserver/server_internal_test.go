package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/sotanodroid/gophiato/internal/app/model"
	"github.com/sotanodroid/gophiato/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticatedUser(t *testing.T) {
	store := teststore.New()
	testUser := model.TestUser(t)
	store.User().Create(testUser)

	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": testUser.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "not authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": 0,
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	s := newServer(store, sessions.NewCookieStore([]byte("secret")))
	sc := securecookie.New([]byte("secret"), nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/", nil)
			cookeStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("cookie", fmt.Sprintf("%s=%s", sessionName, cookeStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleServerusersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"email": "user@example.com",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "invalid_payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("")))
	testUser := model.TestUser(t)
	s.store.User().Create(testUser)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    testUser.Email,
				"password": testUser.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"email":    testUser.Email,
				"password": "wrong",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}
