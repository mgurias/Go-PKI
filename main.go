package main

import (
	"Go-PKI/database"
	"Go-PKI/handlers"
	"log"

	_ "Go-PKI/docs"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {

	if database.TestConn() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	err := database.MongoMigrations()

	if err != nil {
		log.Fatal("Error al aplicar las migraciones")
		return
	}

	handlers.Handlers()

}
