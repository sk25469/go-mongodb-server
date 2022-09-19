package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sk25469/go-mongodb-server/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mg *models.MongoInstance

func Connect(DatabaseUserName string, DatabasePassword string) error {
	const dbName = "go-fiber-hrms"
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@imsahil.prhjfxr.mongodb.net/?retryWrites=true&w=majority", DatabaseUserName, DatabasePassword)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("couldn't connect to database")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = &models.MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func GetMongoInstance() *models.MongoInstance {
	return mg
}
