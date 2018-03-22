package models

import (
	"Ingress/src/db"
	"Ingress/src/validator"
	"log"
	"regexp"
	"strings"
)

//Owner - struct for containing owership for the warehouse
type Owner struct {
	Email string `json:"Email" binding:"required"`
}

// Warehouse - struct for warehouse data exchange between client and server
type Warehouse struct {
	Owner  Owner       `json:"Owner" binding:"required"`
	Name   string      `json:"Name" binding:"required"`
	DBConn *db.Session `json:"-"`
}

func (w *Warehouse) checkIfEmptyRequest() bool {
	if len(strings.TrimSpace(w.Owner.Email)) == 0 || len(strings.TrimSpace(w.Name)) == 0 {
		return true
	}

	return false
}

func (w *Warehouse) checkOwner() bool {
	//TODO: this will different for checking the owner used right now for placeholder
	r, _ := regexp.Compile(`^(("[\w-\s]+")|([\w-]+(?:\.[\w-]+)*)|("[\w-\s]+")([\w-]+(?:\.[\w-]+)*))(@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$)|(@\[?((25[0-5]\.|2[0-4][0-9]\.|1[0-9]{2}\.|[0-9]{1,2}\.))((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\.){2}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\]?$)`)
	return !(r.MatchString(w.Owner.Email))
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

//AddWarehouse - add warehouse to the database
func (w *Warehouse) AddWarehouse() error {
	if err := w.DBConn.DB("ingress").C("warehouses").Insert(w); err != nil {
		//NOTE: probably shouldn't be fatal - or am i dumb and think this closes the program?
		log.Fatalf("Error inserting to db %s", err)
		return err
	}

	return nil
}
