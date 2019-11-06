package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/yota-hada/mongo-go/model"
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

	var questions []interface{}

	for index := 0; index < 5; index++ {
		indexString := strconv.Itoa(index)

		question := model.Question{
			LineChannelID: indexString,
			Title:         fmt.Sprintf("アンケートタイトル%s", indexString),
			Question: []model.QuestionData{
				{
					Title:   "質問1",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問2",
					Answers: []string{"A", "B", "C", "D"},
				},
			},
		}

		questions = append(questions, question)
	}

	res, err := questionColl.InsertMany(context.Background(), questions)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
