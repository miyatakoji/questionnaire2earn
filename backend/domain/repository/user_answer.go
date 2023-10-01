package repository

import "a2e/domain/model"

type UserAnswerRepo interface {
	Answer(model.UserAnswer) error
	GetBySurveyID(SurveyId int64) (model.Question, error)
}
