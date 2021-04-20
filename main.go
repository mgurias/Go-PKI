package main

import (
	"Go-PKI/database"
	"Go-PKI/handlers"
	"log"

	_ "Go-PKI/docs"
)

// @title PKI API
// @version 1.0
// @description Ejemplos de servicios necesarios para administrar la generacion de PKI
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
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
