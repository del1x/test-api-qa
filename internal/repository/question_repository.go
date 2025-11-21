package repository

import (
	"qa-service/internal/model"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(question *model.Question) error
	GetAll() ([]model.Question, error)
	GetByID(id uint) (*model.Question, error)
	Delete(id uint) error
	PreloadAnswers(id uint) (*model.Question, error)
}

type questionRepo struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepo{db}
}

func (r *questionRepo) Create(q *model.Question) error {
	return r.db.Create(q).Error
}

func (r *questionRepo) GetAll() ([]model.Question, error) {
	var questions []model.Question
	err := r.db.Find(&questions).Error
	return questions, err
}

func (r *questionRepo) GetByID(id uint) (*model.Question, error) {
	var q model.Question
	err := r.db.First(&q, id).Error
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *questionRepo) PreloadAnswers(id uint) (*model.Question, error) {
	var q model.Question
	err := r.db.Preload("Answers").First(&q, id).Error
	return &q, err
}

func (r *questionRepo) Delete(id uint) error {
	return r.db.Delete(&model.Question{}, id).Error
}
