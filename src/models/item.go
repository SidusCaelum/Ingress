package models

import (
	"Ingress/src/db"
	"image/png"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

//Item - struct for containing item information
type Item struct {
	Name        string      `json:"Name" binding:"required"`
	Description string      `json:"Description" binding:"required"`
	Warehouse   string      `json:"Warehouse" binding:"required"`
	Date        string      `json:"Date" binding:"required"`
	Time        string      `json:"Time" binding:"required"`
	UUID        string      `json:"UUID" binding:"required"`
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
