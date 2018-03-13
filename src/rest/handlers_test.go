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
		IsEmpty:     false,
		BadUsername: false,
		BadEmail:    false,
	}

	actualResponse := &UserCheck{}

	err = json.NewDecoder(w.Body).Decode(&actualResponse)
	if err != nil {
		t.Errorf("Error when decoding response from NewUser endpoint")
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
	assert.Equal(t, expectedResponse.BadUsername, actualResponse.BadUsername)
	assert.Equal(t, expectedResponse.BadEmail, actualResponse.BadEmail)
}

// TestEmptyNewUserEmailUsername - check if system handles empty submission
func TestEmptyNewUserEmailUsername(t *testing.T) {
	r := NewRouter(true)
	user := &User{Email: "", Username: ""}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error with TestUser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)

	expectedResponse := &UserCheck{
		IsEmpty:     true,
		BadUsername: true,
		BadEmail:    true,
	}

	actualResponse := &UserCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("Error when decoding response from NewUser endpoint")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
}

// testbadnewuseremailusername - check if system handles incorrect submission
func TestBadNewUserEmailUsername(t *testing.T) {
	r := NewRouter(true)
	user := &User{Email: "test%^@testing.$go", Username: "*lkjjsdf*"}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("error with testuser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)

	expectedResponse := &UserCheck{
		IsEmpty:     false,
		BadUsername: true,
		BadEmail:    true,
	}

	actualResponse := &UserCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("error when decoding response from newuser endpoint")
	}

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
	assert.Equal(t, expectedResponse.BadUsername, actualResponse.BadUsername)
	assert.Equal(t, expectedResponse.BadEmail, actualResponse.BadEmail)
}
