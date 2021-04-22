package routes

import (
	"Go-PKI/database"
	"Go-PKI/models"
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor de Email usado en todos los EndPoints */
var Correo string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, exists, _ := database.TestUserExists(claims.Correo)
		if exists {
			Correo = claims.Correo
			IDUsuario = claims.ID.Hex()
		}
		return claims, true, IDUsuario, nil
	}
	if !token.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
