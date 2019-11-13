package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rs/xid"
	"github.com/yota-hada/mongo-go/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SurveyLog is
type SurveyLog struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	EnqueteID  string             `json:"enqueteID" bson:"enqueteID"`
	LineUserID string             `json:"lineUserID" bson:"lineUserID"`
	Result     map[string]string  `json:"result" bson:"result"`
}

// TODO: 関数か
func main() {
	clientCount := flag.Uint("clientCount", 0, "Client Count(>= 0)")
	enqueteCount := flag.Uint("enqueteCount", 0, "Enquete Count(>= 0)")
	lineUserCount := flag.Uint("lineUserCount", 0, "lineUserCount Count(>= 0)")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	surveyColl := client.Database("lycle_line").Collection("survey_log")

	for clientIndex := 0; clientIndex < int(*clientCount); clientIndex++ {
		for enqueteIndex := 0; enqueteIndex < int(*enqueteCount); enqueteIndex++ {
			enqueteID := xid.New().String()
			questionID1 := xid.New().String()
			questionID2 := xid.New().String()
			questionID3 := xid.New().String()
			questionID4 := xid.New().String()
			questionID5 := xid.New().String()
			questionID6 := xid.New().String()
			questionID7 := xid.New().String()
			questionID8 := xid.New().String()
			questionID9 := xid.New().String()
			questionID10 := xid.New().String()

			var surverLogs []interface{}

			for i := 0; i < int(*lineUserCount); i++ {
				lineUserID := xid.New().String()

				// NOTE: bson.Dだとfieldの順番を担保してくれる
				surveyLog := bson.D{
					{Key: "_id", Value: primitive.NewObjectID()},
					{Key: "enqueteID", Value: enqueteID},
					{Key: "lineUserID", Value: lineUserID},
					// NOTE: 回答はA~Eをランダムで生成
					// HACK: ダミーの回答が重複してる点
					{Key: questionID1, Value: bson.A{helper.GenerateAnswer(), helper.GenerateAnswer(), helper.GenerateAnswer()}},
					{Key: questionID2, Value: helper.GenerateAnswer()},
					{Key: questionID3, Value: helper.GenerateAnswer()},
					{Key: questionID4, Value: helper.GenerateAnswer()},
					{Key: questionID5, Value: helper.GenerateAnswer()},
					{Key: questionID6, Value: helper.GenerateAnswer()},
					{Key: questionID7, Value: helper.GenerateAnswer()},
					{Key: questionID8, Value: helper.GenerateAnswer()},
					{Key: questionID9, Value: helper.GenerateAnswer()},
					{Key: questionID10, Value: helper.GenerateAnswer()},
				}

				surverLogs = append(surverLogs, surveyLog)
			}

			res, err := surveyColl.InsertMany(context.Background(), surverLogs)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(res)
		}
	}
}
