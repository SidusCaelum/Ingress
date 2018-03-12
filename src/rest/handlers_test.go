package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, jsonUser []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(jsonUser))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// TestValidNewUserEmailUsername - check if system handles correct submission
func TestValidNewUserEmailUsername(t *testing.T) {
	r := NewRouter(true)
	user := &User{Email: "test@test.com", Username: "test"}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error with TestUser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)

	expectedResponse := &UserCheck{
        Empty: true,
        Username: true,
        Email: true,
    }

	actualResponse := &UserCheck{}

	err = json.NewDecoder(w.Body).Decode(&actualResponse)
	if(err != nil) {
		t.Errorf("Error when decoding response from NewUser endpoint")
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse.Empty, actualResponse.Empty)
	assert.Equal(t, expectedResponse.Username, actualResponse.Username)
	assert.Equal(t, expectedResponse.Email, actualResponse.Email)
}

// TestNonValidNewUserEmailUsername - check if system handles incorrect submission
func TestNonValidNewUserEmailUsername(t *testing.T) {
	r := NewRouter(true)
	user := &User{Email: "", Username: ""}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error with TestUser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)

	expectedResponse := &UserCheck{
		Empty: false,
		Username: false,
		Email: false,
	}

	actualResponse := &UserCheck{}

	err = json.NewDecoder(w.Body).Decode(&actualResponse)
	if err != nil {
		t.Errorf("Error when decoding response from NewUser endpoint")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedResponse.Empty, actualResponse.Empty)
}
