package model

// HACK: クリーンアーキテクチャ的にはエンティティが外部に依存するのはアウト
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: andとor検索は詳細の仕様を詰める

// Segment is
type Segment struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	QuestionID    primitive.ObjectID `json:"questionID" bson:"questionID"`
	LineChannelID string             `json:"lineChannelID" bson:"lineChannelID"`
	SegmentData   []QuestionData     `json:"segmentData" bson:"segmentData"`
	IsOr          bool               `json:"isOr" bson:"isOr"`
}
