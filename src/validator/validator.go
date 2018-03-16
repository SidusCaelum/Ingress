package validator

import (
	"reflect"
)

// Validator - interface for all validator checks
type Validator interface {
	Run() interface{}
}

//Validate - run the checker
func Validate(v Validator) (interface{}, bool) {
	s := v.Run()
	status := true

	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)

		if valueField.Interface() == true {
			//if one element of UserCheck is false status is false return
			status = false
			break
		}
	}

	return s, status
}

// UserCheck - struct for checking user submission content
type UserCheck struct {
	IsEmpty     bool `json:"IsEmpty"`
	BadUsername bool `json:"BadUsername"`
	BadEmail    bool `json:"BadEmail"`
}

// WarehouseCheck - struct for checking warehouse submission content
type WarehouseCheck struct {
	IsEmpty          bool `json:"IsEmpty"`
	BadOwner         bool `json:"BadUsername"`
	BadWarehouseName bool `json:"BadWarehouseName"`
}
