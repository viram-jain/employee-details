package config

import (
	"fmt"
	"os"

	"bitbucket.org/kaleyra/mongo-sdk/mongo"
	"github.com/viramjainkaleyra/employee-details/model"
)

var collection *mongo.Collection

// ConnectToMongo creates a mongodb connection
func ConnectToMongo() {
	db := mongo.URI{
		Username: "",
		Password: "",
		Host:     os.Getenv("MONGO_HOST"),
		DB:       os.Getenv("MONGO_DATABASE"),
		Port:     os.Getenv("MONGO_PORT"),
	}
	client, err := mongo.NewClient(db)
	if err != nil {
		sugarLogger.Errorf("Failed to connect to mongodb = %s", err.Error())
		return
	}
	fmt.Println("Connected to MongoDB!")
	collection = client.Collection(os.Getenv("MONGO_COLLECTION"))
	fmt.Println("Collection instance created!")
}

// AddToMongo adds a single record in the mongodb
func AddToMongo(employee model.Employee) {
	insertResult, err := collection.InsertOne(employee)
	if err != nil {
		sugarLogger.Errorf("Failed to insert data %s", err.Error())
		return
	}
	fmt.Println("Inserted data -> ", insertResult)
}
