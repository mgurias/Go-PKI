package handlers

import (
	"Go-PKI/middlewares"
	"Go-PKI/routes"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv"

	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/Login", middlewares.TestDB(routes.CreateLogin)).Methods("POST")
	router.HandleFunc("/CreateUser", middlewares.TestDB(routes.CreateUser)).Methods("POST")
	router.HandleFunc("/GetUser", middlewares.TestDB(routes.GetUser)).Methods("GET")
	router.HandleFunc("/ModifyUser", middlewares.TestDB(routes.ModifyUser)).Methods("POST")
	router.HandleFunc("/DropUser", middlewares.TestDB(routes.DropUser)).Methods("DELETE")
	//router.HandleFunc("/GetUsers", middlewares.TestDB(routes.GetUsers)).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
