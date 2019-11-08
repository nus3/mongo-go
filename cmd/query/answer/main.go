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
	filter := bson.M{
		"questionID": question.ID,
		"$and": []bson.M{
			bson.M{
				"answers": bson.M{
					"$elemMatch": bson.M{
						"title":  "質問1",
						"answer": "B",
					},
				},
			},
			bson.M{
				"answers": bson.M{
					"$elemMatch": bson.M{
						"title":  "質問2",
						"answer": "A",
					},
				},
			},
		},
	}

	answerAggre, err := answerColl.Find(ctx, filter)
	if err != nil {
		log.Fatalln(err)
	}

	count := 0

	defer answerAggre.Close(ctx)
	for answerAggre.Next(ctx) {
		var result bson.M
		err := answerAggre.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		count++
		fmt.Println(result)
	}
	if err := answerAggre.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
