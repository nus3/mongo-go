package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/yota-hada/mongo-go/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: 関数化
func main() {
	count := flag.Uint("count", 0, "Count(>= 0)")
	flag.Parse()

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

	// TODO: answerをランダムで作成するところから

	answerColl := client.Database("lycle_line").Collection("answer")

	var answers []interface{}

	// HACK: uintとintってどっちに合わせるのがいいんだっけ？
	for index := 0; index < int(*count); index++ {
		indexString := strconv.Itoa(index)

		answer := model.Question{}

		answers = append(answers, answer)
	}

	res, err := questionColl.InsertMany(context.Background(), answers)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
