package model

// Answer is
type Answer struct {
	QuestionID    string       `json:"questionID" bson:"questionID"`
	LineChannelID string       `json:"lineChannelID" bson:"lineChannelID"`
	LineUserID    string       `json:"lineUserID" bson:"lineUserID"`
	Answers       []AnswerData `json:"answers" bson:"answers"`
}

// AnswerData is
type AnswerData struct {
	Title  string `json:"title" bson:"title"`
	Answer string `json:"answer" bson:"answer"`
}
