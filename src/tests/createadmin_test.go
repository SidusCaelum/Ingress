package tests

import (
	"Ingress/src/rest"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	Email    string
	Username string
}

func performRequest(r http.Handler, method, path string, jsonUser []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(jsonUser))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// TestValidNewUser - check if new user can be made
func TestValidNewUser(t *testing.T) {
	r := rest.NewRouter()
	user := &TestUser{Email: "test@test.com", Username: "test"}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error with TestUser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)
	assert.Equal(t, http.StatusOK, w.Code)
	jsonResult, err := json.Marshal(gin.H{"Result": "OK"})
	assert.Equal(t, string(jsonResult)+"\n", w.Body.String())
}
