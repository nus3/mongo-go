package model

// Question is
type Question struct {
	LineChannelID string
	Title         string
	Question      []QuestionData
}

// QuestionData is
type QuestionData struct {
	Title   string
	Answers []string
}
