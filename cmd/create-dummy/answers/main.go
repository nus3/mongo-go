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

// TODO: 関数化
func main() {
	// count := flag.Uint("count", 0, "Count(>= 0)")
	// flag.Parse()

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
	answer := model.Answer{
		QuestionID:    question.ID,
		LineChannelID: question.LineChannelID,
		LineUserID:    "1111111111111",
		// TODO: 回答をランダムにする
		Answers: []model.AnswerData{
			{
				Title:  "質問1",
				Answer: "A",
			},
			{
				Title:  "質問2",
				Answer: "A",
			},
			{
				Title:  "質問3",
				Answer: "A",
			},
			{
				Title:  "質問4",
				Answer: "A",
			},
		},
	}

	res, err := answerColl.InsertOne(context.Background(), answer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)

	// var answers []interface{}

	// // HACK: uintとintってどっちに合わせるのがいいんだっけ？
	// for index := 0; index < int(*count); index++ {
	// 	indexString := strconv.Itoa(index)

	// 	answer := model.Question{}

	// 	answers = append(answers, answer)
	// }

	// res, err := questionColl.InsertMany(context.Background(), answers)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(res)
}
