package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rs/xid"
	"github.com/yota-hada/mongo-go/helper"
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

				surveyLog := SurveyLog{
					ID:         primitive.NewObjectID(),
					EnqueteID:  enqueteID,
					LineUserID: lineUserID,
					Result: map[string]string{
						// NOTE: 回答はA~Eをランダムで生成
						questionID1:  helper.GenerateAnswer(),
						questionID2:  helper.GenerateAnswer(),
						questionID3:  helper.GenerateAnswer(),
						questionID4:  helper.GenerateAnswer(),
						questionID5:  helper.GenerateAnswer(),
						questionID6:  helper.GenerateAnswer(),
						questionID7:  helper.GenerateAnswer(),
						questionID8:  helper.GenerateAnswer(),
						questionID9:  helper.GenerateAnswer(),
						questionID10: helper.GenerateAnswer(),
					},
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
