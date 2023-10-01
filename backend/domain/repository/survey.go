package repository

import "a2e/domain/model"

type SurveyRepo interface {
	GetAll() ([]model.Survey, error)
}
