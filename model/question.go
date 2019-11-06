package model

// HACK: クリーンアーキテクチャ的にはエンティティが外部に依存するのはアウト
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Question is
type Question struct {
	LineChannelID string
	Title         string
	Questions     []QuestionData
}

// QuestionGetRequest is
type QuestionGetRequest struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

// QuestionGetResponse is
type QuestionGetResponse struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	LineChannelID string
	Title         string
	Questions     []QuestionData
}

// QuestionData is
type QuestionData struct {
	Title   string
	Answers []string
}
