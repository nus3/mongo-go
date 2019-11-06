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

	var questions []interface{}

	// HACK: uintとintってどっちに合わせるのがいいんだっけ？
	for index := 0; index < int(*count); index++ {
		indexString := strconv.Itoa(index)

		question := model.Question{
			LineChannelID: indexString,
			Title:         fmt.Sprintf("アンケートタイトル%s", indexString),
			Questions: []model.QuestionData{
				{
					Title:   "質問1",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問2",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問3",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問4",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問5",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問6",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問7",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問8",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問9",
					Answers: []string{"A", "B", "C", "D"},
				},
				{
					Title:   "質問10",
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
