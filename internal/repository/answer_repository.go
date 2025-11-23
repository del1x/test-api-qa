package repository

import (
	"qa-service/internal/model"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(a *model.Answer) error
	GetByID(id uint) (*model.Answer, error)
	Delete(id uint) error
}

type answerRepo struct{ db *gorm.DB }

func NewAnswerRepository(db *gorm.DB) AnswerRepository { return &answerRepo{db} }

func (r *answerRepo) Create(a *model.Answer) error { return r.db.Create(a).Error }
func (r *answerRepo) GetByID(id uint) (*model.Answer, error) {
	var a model.Answer
	return &a, r.db.First(&a, id).Error
}
func (r *answerRepo) Delete(id uint) error { return r.db.Delete(&model.Answer{}, id).Error }
