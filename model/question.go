package model

// HACK: クリーンアーキテクチャ的にはエンティティが外部に依存するのはアウト
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HACK: 色々汚いの
//       IDがあったりなかったり ちゃんとエンティティの設計をする

// Question is
type Question struct {
	LineChannelID string         `json:"lineChannelID" bson:"lineChannelID"`
	Title         string         `json:"title" bson:"title"`
	Questions     []QuestionData `json:"questions" bson:"questions"`
}

// QuestionGetRequest is
type QuestionGetRequest struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

// QuestionGetRequestByLine is
type QuestionGetRequestByLine struct {
	LineChannelID string `json:"lineChannelID" bson:"lineChannelID"`
}

// QuestionGetResponse is
type QuestionGetResponse struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	LineChannelID string             `json:"lineChannelID" bson:"lineChannelID"`
	Title         string             `json:"title" bson:"title"`
	Questions     []QuestionData     `json:"questions" bson:"questions"`
}

// QuestionData is
type QuestionData struct {
	Title   string   `json:"title" bson:"title"`
	Answers []string `json:"answers" bson:"answers"`
}
