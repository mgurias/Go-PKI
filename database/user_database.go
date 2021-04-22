package database

import (
	"Go-PKI/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func TestUserExists(correo string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database(AtlasDatabase)
	col := db.Collection("user")

	condition := bson.M{"correo": correo}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false, ""
	}

	ID := result.ID.Hex()
	result.Password = ""
	return result, true, ID
}

func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

/*InsertUser Inserta los datos de un usuario de la BD */
func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database(AtlasDatabase)
	col := db.Collection("user")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

/*GetUser obtiene los datos de un usuario de la BD */
func SelectUser(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConn.Database(AtlasDatabase)
	col := db.Collection("user")

	var t models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&t)
	t.Password = ""
	if err != nil {
		return t, err
	}
	return t, nil
}

/*UpdateUser permite modificar los datos del usuario en la BD*/
func UpdateUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database(AtlasDatabase)
	col := db.Collection("user")

	row := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		row["nombre"] = u.Nombre
	}
	if len(u.ApellidoPaterno) > 0 {
		row["apellidopaterno"] = u.ApellidoPaterno
	}
	if len(u.ApellidoMaterno) > 0 {
		row["apellidomaterno"] = u.ApellidoMaterno
	}
	if len(u.Curp) > 0 {
		row["curp"] = u.Curp
	}
	if len(u.Rfc) > 0 {
		row["rfc"] = u.Rfc
	}

	updtString := bson.M{
		"$set": row,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, condition, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}

/*DeleteUser Borra los datos de un usuario de la BD */
func DeleteUser(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConn.Database(AtlasDatabase)
	col := db.Collection("user")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	_, err := col.DeleteOne(ctx, condition)

	if err != nil {
		return err
	}
	return nil
}
