package main

import (
	"Go-PKI/routes"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var id = "607f1c7b37a3f258960dca3c"
var json = `{"id":"607f1c7b37a3f258960dca3c","nombre":"Maria","apellidopaterno":"Torres","apellidomaterno":"Arias","correo":"mata@gmail.com","password":"PasswordTemporal","curp":"TOAM991213MSLRSD02","rfc":"TOAM991213MSL"}`
var jsonResult = `{"id":"607f1c7b37a3f258960dca3c","nombre":"Maria","apellidopaterno":"Torres","apellidomaterno":"Arias","correo":"mata@gmail.com","curp":"TOAM991213MSLRSD02","rfc":"TOAM991213MSL"}`
var jsonModify = `{"id":"607f1c7b37a3f258960dca3c","nombre":"Marielena","apellidopaterno":"Torres","apellidomaterno":"Arias","correo":"mata@gmail.com","curp":"TOAM991213MSLRSD02","rfc":"TOAM991213MSL"}`

func TestCreateUser(t *testing.T) {

	var jsonStr = []byte(json)

	req, err := http.NewRequest("POST", "/CreateUser", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/GetUser", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v",
			status, http.StatusOK)
	}

	if strings.TrimRight(rr.Body.String(), "\n") != jsonResult {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), jsonResult)
	}

}
func TestModifyUser(t *testing.T) {
	var jsonStr = []byte(jsonModify)

	req, err := http.NewRequest("PUT", "/ModifyUser", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.ModifyUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDropUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/DropUser", nil)

	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.DropUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v",
			status, http.StatusOK)
	}

}
