package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	ApellidoPaterno string             `bson:"apellidopaterno" json:"apellidopaterno,omitempty"`
	ApellidoMaterno string             `bson:"apellidomaterno" json:"apellidomaterno"`
	Correo          string             `bson:"correo" json:"correo,omitempty"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Curp            string             `bson:"curp" json:"curp,omitempty"`
	Rfc             string             `bson:"rfc" json:"rfc,omitempty"`
}
