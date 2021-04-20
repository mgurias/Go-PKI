package database

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var AtlasUser = os.Getenv("MONGO_ATLAS_USER_NAME")
var AtlasPassword = os.Getenv("MONGO_ATLAS_PASSWORD")
var AtlasHost = os.Getenv("MONGO_ATLAS_HOST")
var AtlasDatabase = os.Getenv("MONGO_ATLAS_DATABASE")
var AtlasOption = os.Getenv("MONGO_ATLAS_OPTIONS")

var MongoConn = ConnDB()

//Connect to database
func ConnDB() *mongo.Client {
	fmt.Println("MONGO_ATLAS_DATABASE: ", AtlasUser)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s%s", AtlasUser, AtlasPassword, AtlasHost, AtlasDatabase, AtlasOption)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	//log.Println("Conexión a la base de datos exitosa")
	return client
}

//Test the connection
func TestConn() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
