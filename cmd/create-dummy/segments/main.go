package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yota-hada/mongo-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: 関数化
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

	segmentColl := client.Database("lycle_line").Collection("segment")

	segment := model.Segment{
		ID:            primitive.NewObjectID(),
		QuestionID:    question.ID,
		LineChannelID: question.LineChannelID,
		SegmentData: []model.QuestionData{
			{
				Title:   "質問1",
				Answers: []string{"A", "C"},
			},
			{
				Title:   "質問2",
				Answers: []string{"B"},
			},
			{
				Title:   "質問3",
				Answers: []string{"D", "A"},
			},
		},
		IsOr: false,
	}

	res, err := segmentColl.InsertOne(context.Background(), segment)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
