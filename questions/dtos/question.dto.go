package dto

import (
	"strings"

	"github.com/isazobu/dailyQuestionsAPI/questions/models"
)

type QuestionDTO struct {
	Course      string `json:"course" form:"title" bson:"course" validate:"required,min=3"`
	Image       string `json:"image" bson:"image" validate:"url"`
	Description string `json:"description" bson:"description" validate:"required"`
	Subject     string `json:"subject" bson:"subject" validate:"required"`
	Difficulty  string `json:"difficulty" bson:"difficulty" validate:"required,oneof=Easy Medium Hard"`

	Option_A string `json:"option_a" bson:"option_a" validate:"required"`
	Option_B string `json:"option_b" bson:"option_b" validate:"required"`
	Option_C string `json:"option_c" bson:"option_c" validate:"required"`
	Option_D string `json:"option_d" bson:"option_d" validate:"required"`
	Option_E string `json:"option_e" bson:"option_e" validate:"required"`
	Answer   string `json:"answer" bson:"answer" validate:"required"`
}

func (q QuestionDTO) MapToQuestionModel() models.Question {
	var target models.Question
	target.Course = q.Course
	target.Image = q.Image
	target.Description = q.Description
	target.Subject = strings.ToLower(q.Subject)
	target.Difficulty = strings.ToLower(q.Difficulty)
	target.Option_A = q.Option_A
	target.Option_B = q.Option_B
	target.Option_C = q.Option_C
	target.Option_D = q.Option_D
	target.Option_E = q.Option_E
	target.Answer = q.Answer
	return target
}

func (q *QuestionDTO) MapFromQuestionModel(src models.Question) {
	q.Course = src.Course
	q.Image = src.Image
	q.Description = src.Description
	q.Subject = src.Subject
	q.Difficulty = src.Difficulty
	q.Option_A = src.Option_A
	q.Option_B = src.Option_B
	q.Option_C = src.Option_C
	q.Option_D = src.Option_D
	q.Option_E = src.Option_E
	q.Answer = src.Answer
}
