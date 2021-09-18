package database

import (
	"context"
	"fmt"
	"time"

	"github.com/calebmchenry/mindful/web-server/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	*mongo.Client
}

var DB *mongo.Client

func Init() *mongo.Client {
	if DB != nil {
		return DB
	}
	username := env.GetMongoUser()
	password := env.GetMongoPassword()
	clusterAddress := "cluster0.dhb5w"
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", username, password, clusterAddress)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DB, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return DB
}

func TestInit() *mongo.Client {
	if DB != nil {
		return DB
	}
	DB = &mongo.Client{}
	return nil
}

func Get() *mongo.Client {
	return DB
}
