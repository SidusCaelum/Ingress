package models

import (
	"Ingress/src/db"
	"Ingress/src/validator"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Item - struct for containing item information
type Item struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Warehouse   string      `json:"warehouse" binding:"required"`
	Date        string      `json:"date" binding:"required"`
	Time        string      `json:"time" binding:"required"`
	UUID        string      `json:"uuid" binding:"required"`
	DBConn      *db.Session `json:"-"`
}

func (i *Item) checkIfEmptyRequest() bool {
	if len(strings.TrimSpace(i.UUID)) == 0 ||
		len(strings.TrimSpace(i.Name)) == 0 ||
		len(strings.TrimSpace(i.Description)) == 0 {
		return true
	}

	return false
}

func (i *Item) checkWarehouse() bool {
	collections, err := i.DBConn.DB(db.DB_NAME).CollectionNames()
	if err != nil {
		log.Printf("Failed to get collection names: %s\n", err)
		return true
	}

	for _, name := range collections {
		if name == i.Warehouse {
			break
		}
	}

	return false
}

func (i *Item) checkUUID() bool {
	//NOTE: setting warehouse here from the item might not be best
	//although this will probably be handled by the return to the client by polling
	//the db first so your selection can only be based off what already exists\
	type Result struct {
		UUID bool `json:"uuid"`
	}

	//NOTE: this query might not be that great might want to revisit
	result := Result{}
	c := i.DBConn.DB(db.DB_NAME).C(i.Warehouse)
	if err := c.Find(nil).Select(bson.M{"uuid": i.UUID}).One(&result); err != nil {
		log.Printf("CHECK UUID: %s\n", err)
		return true
	}

	return false
}

// Run - generic for returning check on Item model
func (i *Item) Run() interface{} {
	itemCheck := &validator.ItemCheck{
		IsEmpty:      i.checkIfEmptyRequest(),
		BadWarehouse: i.checkWarehouse(),
		BadUUID:      i.checkUUID(),
	}

	return itemCheck
}

//AddItem - add item to the database
func (i *Item) AddItem() (bool, error) {
	var err error

	c := i.DBConn.DB(db.DB_NAME).C(i.Warehouse)
	index := mgo.Index{
		Key:    []string{"uuid"},
		Unique: true,
	}

	if err = c.EnsureIndex(index); err != nil {
		//TODO: Come back here and handle this nicely
		panic(err)
	}

	if err = c.Insert(i.marshalJSON()); err != nil {
		if !mgo.IsDup(err) {
			//NOTE: probably shouldn't be fatal - or am i dumb and think this closes the program?
			log.Printf("Error inserting Item to db: %s\n", err)
			return false, err
		}
	}

	//NOTE: this might not need to be here
	return false, nil
}

//ToQR - returns qr code
func (i *Item) ToQR() []byte {
	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
	return nil
}

//HACK: Needs to be done better? mgo continues to serialize DBConn with json:"-"
func (i *Item) marshalJSON() (interface{}, error) {
	var tmp struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		Warehouse   string `json:"warehouse" binding:"required"`
		Date        string `json:"date" binding:"required"`
		Time        string `json:"time" binding:"required"`
		UUID        string `json:"uuid" binding:"required"`
	}

	tmp.Name = i.Name
	tmp.Description = i.Description
	tmp.Warehouse = i.Warehouse
	tmp.Date = i.Date
	tmp.Time = i.Time
	tmp.UUID = i.UUID

	return tmp, nil
}
