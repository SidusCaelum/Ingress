package models

import (
	"Ingress/src/validator"
	"regexp"
	"strings"
)

// Warehouse - struct for warehouse data exchange between client and server
type Warehouse struct {
	Owner string `json:"Owner" binding:"required"`
	Name  string `json:"Name" binding:"required"`
}

func (w *Warehouse) checkIfEmptyRequest() bool {
	if len(strings.TrimSpace(w.Owner)) == 0 || len(strings.TrimSpace(w.Name)) == 0 {
		return true
	}

	return false
}

func (w *Warehouse) checkOwner() bool {
	//TODO: this will different for checking the owner used right now for placeholder
	r, _ := regexp.Compile(`^([a-zA-Z]+[,.]?[ ]?|[a-z]+['-]?)+$`)
	return !(r.MatchString(w.Owner))
}

func (w *Warehouse) checkWarehouseName() bool {
	//TODO: Change regex to something that only accepts letters upper and lower with only underscore
	r, _ := regexp.Compile(`^[a-zA-Z0-9_]*$`)
	return !(r.MatchString(w.Name))
}

//Run - generic for returning checker on User model
func (w *Warehouse) Run() interface{} {
	warehouseCheck := &validator.WarehouseCheck{
		IsEmpty:          w.checkIfEmptyRequest(),
		BadOwner:         w.checkOwner(),
		BadWarehouseName: w.checkWarehouseName(),
	}

	return warehouseCheck
}
