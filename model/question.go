package model

// HACK: クリーンアーキテクチャ的にはエンティティが外部に依存するのはアウト
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Question is
type Question struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	QuestionnaireID primitive.ObjectID `json:"questionnaireID" bson:"questionnaireID"`
	LineChannelID   string             `json:"lineChannelID" bson:"lineChannelID"`
	Title           string             `json:"title" bson:"title"`
	Options         []string           `json:"options" bson:"options"`
}
