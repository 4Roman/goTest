package db

import (
	"context"
	"github.com/gusleein/golog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Collection *mongo.Collection

func Init(ctx context.Context) {
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	// Проверяем коннект
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Создаём коллекцию
	Client = client
	Collection = Client.Database("myDb").Collection("UserInfo")
}
