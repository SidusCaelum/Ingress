package models

import (
	"Ingress/src/db"
	"Ingress/src/validator"
	"log"
	"regexp"
	"strings"

	"github.com/globalsign/mgo/bson"
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
	type Result struct {
		Admin bool `json:"admin"`
	}

	//NOTE: Checking if the the user exists and if they have admin rights
	result := Result{}
	c := w.DBConn.DB("ingress").C("users")
	if err := c.Find(bson.M{"email": w.Owner.Email}).Select(bson.M{"admin": 1}).One(&result); err != nil {
		return true
	}

	return false
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
	var err error

	c := w.DBConn.DB("ingress").C(w.Name)
	if err = c.Insert(w.marshalJSON()); err != nil {
		//NOTE: probably shouldn't be fatal - or am i dumb and think this closes the program?
		log.Printf("Error inserting to db %s", err)
		return err
	}

	return nil
}

//HACK: needs to be done better? mgo continues to serialize DBConn with json:"-"
func (w *Warehouse) marshalJSON() interface{} {
	var tmp struct {
		Owner Owner  `json:"Owner" binding:"required"`
		Name  string `json:"Name" binding:"required"`
	}

	tmp.Owner = w.Owner
	tmp.Name = w.Name
	return tmp
}
