package repository

import "a2e/domain/model"

type QuestionRepo interface {
	GetAll() ([]model.Question, error)
	GetBySurveyID(SurveyId int64) (model.Question, error)
}
