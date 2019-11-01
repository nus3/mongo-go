package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	collection := client.Database("test").Collection("hoge")

	bsonD := bson.D{
		{Key: "str1", Value: "abc"},
		{Key: "num1", Value: 1},
		{Key: "str2", Value: "xyz"},
		{Key: "num2", Value: bson.A{2, 3, 4}},
		{Key: "subdoc", Value: bson.D{{Key: "str", Value: "subdoc"}, {Key: "num", Value: 987}}},
		{Key: "date", Value: time.Now()},
	}

	res, err := collection.InsertOne(context.Background(), bsonD)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
