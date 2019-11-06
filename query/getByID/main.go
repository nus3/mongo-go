package main

import (
	"context"
	"log"
	"time"

	"github.com/yota-hada/mongo-go/model"
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
	collection := client.Database("lycle_line").Collection("question")

	var doc model.QuestionGetResponse
	err = collection.FindOne(context.Background(), bson.D{}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		log.Println("Documents not found")
	} else if err != nil {
		log.Fatalln(err)
	}

	req := model.QuestionGetRequest{
		ID: doc.ID,
	}

	var res model.QuestionGetResponse
	err = collection.FindOne(context.Background(), req).Decode(&res)
	if err == mongo.ErrNoDocuments {
		log.Println("Documents not found")
	} else if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)
}
