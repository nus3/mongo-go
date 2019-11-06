package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yota-hada/mongo-go/model"
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

	questionColl := client.Database("lycle_line").Collection("question")

	question := model.Question{
		LineChannelID: "11111111111",
		Title:         "アンケートタイトル1",
		Question: []model.QuestionData{
			{
				Title:   "質問1",
				Answers: []string{"A", "B", "C", "D"},
			},
			{
				Title:   "質問2",
				Answers: []string{"A", "B", "C", "D"},
			},
		},
	}

	res, err := questionColl.InsertOne(context.Background(), question)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
