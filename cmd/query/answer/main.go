package main

import (
	"context"
	"fmt"
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

	questionColl := client.Database("lycle_line").Collection("question")
	lineChannelReq := model.QuestionGetRequestByLine{
		LineChannelID: "2",
	}
	var question model.QuestionGetResponse
	err = questionColl.FindOne(context.Background(), lineChannelReq).Decode(&question)
	if err == mongo.ErrNoDocuments {
		log.Println("Documents not found")
	} else if err != nil {
		log.Fatalln(err)
	}

	answerColl := client.Database("lycle_line").Collection("answer")
	// TODO: andとor検索の実装
	aggre := model.AnswerAggregater{
		QuestionID: question.ID,
		Answers: model.AnswerData{
			Title:  "質問1",
			Answer: "B",
		},
	}

	answerAggre, err := answerColl.Find(ctx, aggre)
	if err != nil {
		log.Fatalln(err)
	}

	defer answerAggre.Close(ctx)
	for answerAggre.Next(ctx) {
		var result bson.M
		err := answerAggre.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	if err := answerAggre.Err(); err != nil {
		log.Fatal(err)
	}
}
