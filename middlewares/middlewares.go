package middlewares

import (
	"net/http"

	"Go-PKI/database"
	"Go-PKI/routes"
)

func TestDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.TestConn() == 0 {
			http.Error(w, "Se perdió la conexión a la base de datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}

/*VerifyJWT permite validar el JWT que nos viene en la petición */
func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
