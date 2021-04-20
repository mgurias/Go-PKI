package models

/*RespuestaLogin tiene el token que se devuelve con el login */
type Token struct {
	Token string `json:"token,omitempty"`
}
