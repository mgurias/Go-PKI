package database

import (
	"context"
	"time"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoMigrations() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database(AtlasDatabase)

	m := migrate.NewMigrate(db, migrate.Migration{
		Version:     1,
		Description: "add user-collection",
		Up: func(db *mongo.Database) error {

			pass, err := EncryptPassword("P@sswordTemporal")
			if err != nil {
				return err
			}

			_, err = db.Collection("user").InsertOne(ctx, bson.D{
				{Key: "nombre", Value: "Marla Genesis"},
				{Key: "apellidopaterno", Value: "Urenda"},
				{Key: "apellidomaterno", Value: "Castillo"},
				{Key: "correo", Value: "mguc@gmail.com"},
				{Key: "password", Value: pass},
				{Key: "curp", Value: "UICG991213MSLRSD02"},
				{Key: "rfc", Value: "UICG991213MSL"},
				//{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
			})

			if err != nil {
				return err
			}

			opt := options.Index().SetName("user-index")
			keys := bson.D{{Key: "correo", Value: 1}}
			model := mongo.IndexModel{Keys: keys, Options: opt}
			_, err = db.Collection("user").Indexes().CreateOne(context.TODO(), model)
			if err != nil {
				return err
			}

			pass, err = EncryptPassword("PasswordTempor@l")
			if err != nil {
				return err
			}

			_, err = db.Collection("user").InsertOne(ctx, bson.D{
				{Key: "nombre", Value: "Humberto"},
				{Key: "apellidopaterno", Value: "Miranda"},
				{Key: "apellidomaterno", Value: "Perez"},
				{Key: "correo", Value: "hump@gmail.com"},
				{Key: "password", Value: pass},
				{Key: "curp", Value: "MIPH991213MSLRSD02"},
				{Key: "rfc", Value: "MIPH991213MSL"},
				//{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
			})

			if err != nil {
				return err
			}

			return nil
		},
		Down: func(db *mongo.Database) error {
			_, err := db.Collection("user").Indexes().DropOne(context.TODO(), "user-index")

			if err != nil {
				return err
			}

			err = db.Collection("user").Drop(context.TODO())

			if err != nil {
				return err
			}

			return nil
		},
	})

	if err := m.Up(migrate.AllAvailable); err != nil {
		return err
	}

	return nil
}
