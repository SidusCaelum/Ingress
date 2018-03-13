package rest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNewWarehouse(t *testing.T) {
	r := NewRouter(true)
	warehouse := &Warehouse{
		Owner: "joshua",
		Name:  "fun_new_Warehouse",
	}

	jsonWarehouse, err := json.Marshal(warehouse)
	if err != nil {
		t.Errorf("error with testwarehouse struct marshalling: %s", err)
	}

	w := performRequest(r, "POST", "/NewWarehouse", jsonWarehouse)

	expectedResponse := &WarehouseCheck{
		IsEmpty:          false,
		BadOwner:         false,
		BadWarehouseName: false,
	}

	actualResponse := &WarehouseCheck{}

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
	r := NewRouter(true)
	warehouse := &Warehouse{
		Owner: "",
		Name:  "",
	}

	jsonWarehouse, err := json.Marshal(warehouse)
	if err != nil {
		t.Errorf("Unable to marshal warehouse struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewWarehouse", jsonWarehouse)

	expectedResponse := &WarehouseCheck{
		IsEmpty:          true,
		BadOwner:         true,
		BadWarehouseName: true,
	}

	actualResponse := &WarehouseCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("Unable to marshal warehouse struct: %s", err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
}

// TestBadNewWarehouse - check if system handles incorrect submission
func TestBadNewWarehouse(t *testing.T) {
	r := NewRouter(true)
	warehouse := &Warehouse{
		Owner: "092394fjj+_+(&&f)",
		Name:  "32984_jsdfj*)(D)",
	}

	jsonWarehouse, err := json.Marshal(warehouse)
	if err != nil {
		t.Errorf("Unable to marshal warehouse struct: %s", err)
	}

	w := performRequest(r, "POST", "/NewWarehouse", jsonWarehouse)

	expectedResponse := &WarehouseCheck{
		IsEmpty:          false,
		BadOwner:         true,
		BadWarehouseName: true,
	}

	actualResponse := &WarehouseCheck{}

	if err = json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
		t.Errorf("Error when decoding response from NewWarehouse endpoint")
	}

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, expectedResponse.IsEmpty, actualResponse.IsEmpty)
	assert.Equal(t, expectedResponse.BadOwner, actualResponse.BadOwner)
	assert.Equal(t, expectedResponse.BadWarehouseName, actualResponse.BadWarehouseName)
}
