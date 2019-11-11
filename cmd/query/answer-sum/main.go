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
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	questionColl := client.Database("lycle_line").Collection("question")
	questionReq := bson.M{
		"lineChannelID": "5dc925e0dc65a2155a3a84c9",
		// "lineChannelID": "5dc8fcf253efeed654ad0b8a",
	}

	answerColl := client.Database("lycle_line").Collection("answer")

	answerSum := map[string][]model.AnswerData{}

	questions, err := questionColl.Find(ctx, questionReq)
	if err != nil {
		log.Fatal(err)
	}
	defer questions.Close(ctx)
	for questions.Next(ctx) {
		var question model.Question
		err := questions.Decode(&question)
		if err != nil {
			log.Fatal(err)
		}

		pipeline := []bson.M{
			bson.M{
				"$match": bson.M{
					"questionID": question.ID,
				},
			},
			bson.M{
				"$group": bson.M{
					"_id": bson.M{
						"questionID": "$questionID",
						"answer":     "$answer",
					},
					"sum": bson.M{
						"$sum": 1,
					},
				},
			},
		}

		answerAggre, err := answerColl.Aggregate(ctx, pipeline)
		if err != nil {
			log.Fatalln(err)
		}

		defer answerAggre.Close(ctx)

		for answerAggre.Next(ctx) {
			var result model.AnswerAggregater
			err := answerAggre.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}

			answerData := model.AnswerData{
				Answer: result.ID.Answer,
				Count:  result.Sum,
			}

			answerSum[result.ID.QuestionID.Hex()] = append(answerSum[result.ID.QuestionID.Hex()], answerData)
		}
		if err := answerAggre.Err(); err != nil {
			log.Fatal(err)
		}
	}
	if err := questions.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(answerSum)

	// 処理
	end := time.Now()
	fmt.Printf("%f秒かかった\n", (end.Sub(start)).Seconds())
}
