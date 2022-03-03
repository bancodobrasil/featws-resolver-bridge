package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bancodobrasil/featws-resolver-bridge/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {

	cfg := config.GetConfig()
	if cfg == nil {
		log.Fatalf("Não foi carregada configuracão!\n")
		return
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, calcel := context.WithTimeout(context.Background(), 10*time.Second)
	defer calcel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	DB = client
}

//Client instace
var DB *mongo.Client

//getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	cfg := config.GetConfig()
	if cfg == nil {
		log.Fatalf("Não foi carregada configuracão!\n")
		return nil
	}
	collection := DB.Database(cfg.MongoDB).Collection(collectionName)
	return collection
}
