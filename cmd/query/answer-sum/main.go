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

	// TODO: piplineの構造体化
	pipeline := []bson.M{
		// bson.M{
		// 	"$match": bson.M{
		// 		"answers.title": "質問1",
		// 	},
		// 	"$sum": "$answers.answer",
		// },
		// bson.M{
		// 	"$group": bson.M{
		// 		"_id": "$answers.title",
		// 		"answer": bson.M{
		// 			"$push": bson.M{
		// 				"answer": "$answers.answer",
		// 			},
		// 		},
		// 	},
		// },
		bson.M{
			"$addFields": bson.M{
				"test": bson.M{
					"$objectToArray": "$answers.answer",
				},
			},
		},
		// bson.M{
		// 	"$count": "$answers.answer",
		// },
	}
	// pipeline := []bson.M{
	// 	bson.M{
	// 		"$match": bson.M{
	// 			"answers": bson.M{
	// 				"title":  "質問1",
	// 				"answer": "B",
	// 			},
	// 		},
	// 	},
	// 	bson.M{
	// 		"$count": "sum",
	// 	},
	// }
	answerAggre, err := answerColl.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatalln(err)
	}

	defer answerAggre.Close(ctx)

	// HACK: for文回したくない　合計数とるのに
	for answerAggre.Next(ctx) {
		var result bson.M
		err := answerAggre.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		break
	}
	if err := answerAggre.Err(); err != nil {
		log.Fatal(err)
	}
}
