package model

type Question struct {
	ID              int64
	QuestionnaireID int64
	Order           int64
	IsMust          bool
	QuestionType    AnswerType
	Choices         []string
	ImageURL        string
}
