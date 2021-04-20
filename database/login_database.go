package database

import (
	"Go-PKI/models"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin verifica si existe el correo en la BD */
func TryLogin(correo string, password string) (models.User, bool) {
	user, ok, _ := TestUserExists(correo)
	if !(ok) {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}

/*CreateJWT genera el encriptado con JWT */
func CreateJWT(t models.User) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))

	payload := jwt.MapClaims{
		"correo":          t.Correo,
		"nombre":          t.Nombre,
		"apellidopaterno": t.ApellidoPaterno,
		"apellidomaterno": t.ApellidoMaterno,
		"curp":            t.Curp,
		"rfc":             t.Rfc,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
