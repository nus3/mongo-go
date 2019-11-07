package model

// HACK: クリーンアーキテクチャ的にはエンティティが外部に依存するのはアウト
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Answer is
type Answer struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	QuestionID    primitive.ObjectID `json:"questionID" bson:"questionID"`
	LineChannelID string             `json:"lineChannelID" bson:"lineChannelID"`
	LineUserID    string             `json:"lineUserID" bson:"lineUserID"`
	Answers       []AnswerData       `json:"answers" bson:"answers"`
}

// AnswerData is
type AnswerData struct {
	Title  string `json:"title" bson:"title"`
	Answer string `json:"answer" bson:"answer"`
}
