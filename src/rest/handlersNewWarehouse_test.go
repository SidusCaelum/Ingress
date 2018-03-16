package rest

import (
	"Ingress/src/db"
	"Ingress/src/models"
	"Ingress/src/validator"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewWarehouse(t *testing.T) {
	//NOTE: Do actually want to do this?
	//Maybe come back to this and think about it more.
	db, err := db.InitDB("localhost")
	if err != nil {
		log.Fatalf("Database connection failed %s", err)
	}

	r := NewRouter(true, db)
	warehouse := &models.Warehouse{
		Owner: "joshua",
		Name:  "fun_new_Warehouse",
	}

	jsonWarehouse, err := json.Marshal(warehouse)
	if err != nil {
		t.Errorf("error with testwarehouse struct marshalling: %s", err)
	}

	w := performRequest(r, "POST", "/NewWarehouse", jsonWarehouse)

	expectedResponse := &validator.WarehouseCheck{
		IsEmpty:          false,
		BadOwner:         false,
		BadWarehouseName: false,
	}

	actualResponse := &validator.WarehouseCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("error when decoding response from newwarehouse endpoint")
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
	assert.Equal(t, expectedResponse.BadOwner, actualResponse.BadOwner)
	assert.Equal(t, expectedResponse.BadWarehouseName, actualResponse.BadWarehouseName)
}

//TestEmptyNewWarehouse - check if system handles empty submission
func TestEmptyNewWarehouse(t *testing.T) {
	db, err := db.InitDB("localhost")
	if err != nil {
		log.Fatalf("Database connection failed %s", err)
	}

	r := NewRouter(true, db)
	warehouse := &models.Warehouse{
		Owner: "",
		Name:  "",
	}

	jsonWarehouse, err := json.Marshal(warehouse)
	if err != nil {
		t.Errorf("Unable to marshal warehouse struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewWarehouse", jsonWarehouse)

	expectedResponse := &validator.WarehouseCheck{
		IsEmpty:          true,
		BadOwner:         true,
		BadWarehouseName: true,
	}

	actualResponse := &validator.WarehouseCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("Unable to marshal warehouse struct: %s", err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
}

// TestBadNewWarehouse - check if system handles incorrect submission
func TestBadNewWarehouse(t *testing.T) {
	db, err := db.InitDB("localhost")
	if err != nil {
		log.Fatalf("Database connection failed %s", err)
	}

	r := NewRouter(true, db)
	warehouse := &models.Warehouse{
		Owner: "092394fjj+_+(&&f)",
		Name:  "32984_jsdfj*)(D)",
	}

	jsonWarehouse, err := json.Marshal(warehouse)
	if err != nil {
		t.Errorf("Unable to marshal warehouse struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewWarehouse", jsonWarehouse)

	expectedResponse := &validator.WarehouseCheck{
		IsEmpty:          false,
		BadOwner:         true,
		BadWarehouseName: true,
	}

	actualResponse := &validator.WarehouseCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("Error when decoding response from NewWarehouse endpoint")
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
	assert.Equal(t, expectedResponse.BadOwner, actualResponse.BadOwner)
	assert.Equal(t, expectedResponse.BadWarehouseName, actualResponse.BadWarehouseName)
}
