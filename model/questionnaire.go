package model

// HACK: クリーンアーキテクチャ的にはエンティティが外部に依存するのはアウト
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Questionnaire is
type Questionnaire struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	LineChannelID string             `json:"lineChannelID" bson:"lineChannelID"`
	Title         string             `json:"title" bson:"title"`
}
