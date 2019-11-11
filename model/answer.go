package model

// HACK: クリーンアーキテクチャ的にはエンティティが外部に依存するのはアウト
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Answer is
type Answer struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	QuestionnaireID primitive.ObjectID `json:"questionnaireID" bson:"questionnaireID"`
	QuestionID      primitive.ObjectID `json:"questionID" bson:"questionID"`
	LineChannelID   string             `json:"lineChannelID" bson:"lineChannelID"`
	LineUserID      string             `json:"lineUserID" bson:"lineUserID"`
	Answer          string             `json:"answer" bson:"answer"`
}

// HACK: もっといい感じのクエリ書いたら構造体をAggregaterとAnswerSumで分けずにすみそう

// AnswerAggregater is
type AnswerAggregater struct {
	ID struct {
		Answer     string             `json:"answer" bson:"answer"`
		QuestionID primitive.ObjectID `json:"questionID" bson:"questionID"`
	} `json:"id" bson:"_id"`
	Sum int `json:"sum" bson:"sum"`
}

// AnswerData is
type AnswerData struct {
	Answer string `json:"answer" bson:"answer"`
	Count  int    `json:"count" bson:"count"`
}
