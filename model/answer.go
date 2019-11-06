package model

// Answer is
type Answer struct {
	QuestionID    string
	LineChannelID string
	LineUserID    string
	Answers       []AnswerData
}

// AnswerData is
type AnswerData struct {
	Title  string
	Answer string
}
