package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/yota-hada/mongo-go/helper"
	"github.com/yota-hada/mongo-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: 関数か
func main() {
	start := time.Now()

	questionCount := flag.Uint("questionCount", 0, "Question Count(>= 0)")
	answerCount := flag.Uint("answerCount", 0, "Answer Count(>= 0)")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	questionnaireColl := client.Database("lycle_line").Collection("questionnaire")

	questionnaireID := primitive.NewObjectID()

	questionnaire := model.Questionnaire{
		ID: questionnaireID,
		// NOTE: 本来はうちのDBに保存されてる物を使う
		LineChannelID: primitive.NewObjectID().Hex(),
		Title:         fmt.Sprintf("アンケート名%s", questionnaireID.Hex()),
	}

	questionnaireRes, err := questionnaireColl.InsertOne(context.Background(), questionnaire)
	if err != nil {
		log.Fatalln(err)
	}

	questionColl := client.Database("lycle_line").Collection("question")

	var questions []interface{}
	var answers []interface{}
	for i := 0; i < int(*questionCount); i++ {
		iString := strconv.Itoa(i + 1)

		question := model.Question{
			ID:              primitive.NewObjectID(),
			QuestionnaireID: questionnaire.ID,
			LineChannelID:   questionnaire.LineChannelID,
			Title:           fmt.Sprintf("問題%s", iString),
			Options: []string{
				"A",
				"B",
				"C",
				"D",
			},
		}

		for j := 0; j < int(*answerCount); j++ {
			jString := strconv.Itoa(j + 1)

			answer := model.Answer{
				ID:              primitive.NewObjectID(),
				QuestionnaireID: questionnaire.ID,
				QuestionID:      question.ID,
				LineChannelID:   questionnaire.LineChannelID,
				LineUserID:      fmt.Sprintf("line-user-%s", jString),
				Answer:          helper.GenerateAnswer(),
			}

			answers = append(answers, answer)
		}

		questions = append(questions, question)
	}

	questionsRes, err := questionColl.InsertMany(context.Background(), questions)
	if err != nil {
		log.Fatalln(err)
	}

	answerColl := client.Database("lycle_line").Collection("answer")
	answerRes, err := answerColl.InsertMany(context.Background(), answers)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(questionnaireRes)
	fmt.Println(questionsRes)
	fmt.Println(answerRes)

	// 処理
	end := time.Now()
	fmt.Printf("%f秒かかった\n", (end.Sub(start)).Seconds())
}
