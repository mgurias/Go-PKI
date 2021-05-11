package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"Go-PKI/database"
	"Go-PKI/models"
)

// CreateLogin godoc
// @Summary Create a new login to database
// @Description Create a new login
// @Tags Login
// @Accept json
// @Produce json
// @Param Login body string true "User/Password"
// @Success 200 {object} models.User
// @Router /Login [post]
func CreateLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos de usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(t.Correo) == 0 {
		http.Error(w, "El correo electrónico del usuario es requerido ", 400)
		return
	}

	document, ok := database.TryLogin(t.Correo, t.Password)
	if ok {
		http.Error(w, "Usuario y/o contraseña inválidos ", 400)
		return
	}

	jwtKey, err := database.CreateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+err.Error(), 400)
		return
	}

	result := models.Token{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

	//Only cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "Datos de usuario"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /CreateUser [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	if len(t.Correo) == 0 {
		http.Error(w, "La cuenta de correo es requerida", 400)
		return
	}
	if len(t.Password) < 8 {
		http.Error(w, "La longitud de la contraseña debe ser al menos de ocho caracteres", 400)
		return
	}

	_, exists, _ := database.TestUserExists(t.Correo)
	if exists {
		http.Error(w, "La cuenta de correo ya existe", 400)
		return
	}

	_, ok, err := database.InsertUser(t)
	if err != nil {
		http.Error(w, "Error al insertar el registro"+err.Error(), 400)
		return
	}
	if !(ok) {
		http.Error(w, "Error al registrar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// GetUser godoc
// @Summary Get data user
// @Description Select data user from DB
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /GetUser [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	t, err := database.SelectUser(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al consultar el ID del usuario "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

// modifyUser godoc
// @Summary Modify user data
// @Description Modify user data from DB
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Param user body models.User true "Datos de Usuario"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /ModifyUser [post]
func ModifyUser(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return

	}

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var ok bool

	ok, err = database.UpdateUser(t, ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if !(ok) {
		http.Error(w, "Ocurrió un error al modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// DropUser godoc
// @Summary Delete data user
// @Description Delete data user from DB
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /DropUser [delete]
func DropUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := database.DeleteUser(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al consultar el ID del ususario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusOK)

}
