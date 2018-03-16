package rest

import (
	"Ingress/src/db"
	"Ingress/src/models"
	"Ingress/src/validator"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TODO: move this out to another file? maybe extract test creation (at least in part the repeated steps)
func performRequest(r http.Handler, method, path string, jsonUser []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(jsonUser))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// TestValidNewUser - check if system handles correct submission
func TestValidNewUser(t *testing.T) {
	db, err := db.InitDB("localhost")
	if err != nil {
		log.Fatalf("Database connection failed %s", err)
	}

	r := NewRouter(true, db)
	user := &models.User{Email: "test@test.com", Username: "test"}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error with TestUser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)

	expectedResponse := &validator.UserCheck{
		IsEmpty:     false,
		BadUsername: false,
		BadEmail:    false,
	}

	actualResponse := &validator.UserCheck{}

	err = json.NewDecoder(w.Body).Decode(&actualResponse)
	if err != nil {
		t.Errorf("Error when decoding response from NewUser endpoint")
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
	assert.Equal(t, expectedResponse.BadUsername, actualResponse.BadUsername)
	assert.Equal(t, expectedResponse.BadEmail, actualResponse.BadEmail)
}

// TestEmptyNewUser - check if system handles empty submission
func TestEmptyNewUser(t *testing.T) {
	db, err := db.InitDB("localhost")
	if err != nil {
		log.Fatalf("Database connection failed %s", err)
	}

	r := NewRouter(true, db)
	user := &models.User{Email: "", Username: ""}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error with TestUser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)

	expectedResponse := &validator.UserCheck{
		IsEmpty:     true,
		BadUsername: true,
		BadEmail:    true,
	}

	actualResponse := &validator.UserCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("Error when decoding response from NewUser endpoint")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
}

// TestBadNewUser - check if system handles incorrect submission
func TestBadNewUser(t *testing.T) {
	db, err := db.InitDB("localhost")
	if err != nil {
		log.Fatalf("Database connection failed %s", err)
	}

	r := NewRouter(true, db)
	user := &models.User{
		Email:    "test%^@testing.$go",
		Username: "*lkjjsdf*",
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Errorf("error with testuser struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewUser", jsonUser)

	expectedResponse := &validator.UserCheck{
		IsEmpty:     false,
		BadUsername: true,
		BadEmail:    true,
	}

	actualResponse := &validator.UserCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("error when decoding response from newuser endpoint")
	}

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
	assert.Equal(t, expectedResponse.BadUsername, actualResponse.BadUsername)
	assert.Equal(t, expectedResponse.BadEmail, actualResponse.BadEmail)
}
