package app

import (
	"Ingress/src/db"
	"Ingress/src/models"
	"Ingress/src/rest"
	"log"

	"github.com/gin-gonic/gin"
)

//App - struct initalizer for ingress
type App struct {
	Router *gin.Engine
	DB     *db.Session
	IP     string
	Port   int
}

//Initialize - initalizer for ingress
func (a *App) Initialize(config *models.StartupConfiguration) {
	a.IP = config.DBIp
	a.Port = config.Port
}

//Run - run app
func (a *App) Run() {
	//Note: This might not work nil pointer perhaps?
	var err error
	a.DB, err = db.InitDB(a.IP)
	if err != nil {
		log.Fatalf("DB connection failed: %s", err)
	}

	a.Router = rest.NewRouter(false, a.DB)

	defer a.DB.Close()
	a.Router.Run()
}
