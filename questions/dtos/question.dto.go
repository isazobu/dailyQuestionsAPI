package dto

import (
	"strings"

	"github.com/isazobu/dailyQuestionsAPI/questions/models"
)

type QuestionDTO struct {
	Title       string `json:"title" form:"title" validate:"required" bson:"title"`
	Image       string `json:"image" bson:"image"`
	Description string `json:"description" bson:"description"`
	Category    string `json:"category" bson:"category"`
	Difficulty  string `json:"difficulty" bson:"difficulty"`

	Option_A string `json:"option_a" bson:"option_a"`
	Option_B string `json:"option_b" bson:"option_b"`
	Option_C string `json:"option_c" bson:"option_c"`
	Option_D string `json:"option_d" bson:"option_d"`
	Option_E string `json:"option_e" bson:"option_e"`
	Answer   string `json:"answer" bson:"answer"`
}

func (q QuestionDTO) MapToQuestionModel() models.Question {
	var target models.Question
	target.Title = q.Title
	target.Image = q.Image
	target.Description = q.Description
	target.Category = strings.ToLower(q.Category)
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
	q.Title = src.Title
	q.Image = src.Image
	q.Description = src.Description
	q.Category = src.Category
	q.Difficulty = src.Difficulty
	q.Option_A = src.Option_A
	q.Option_B = src.Option_B
	q.Option_C = src.Option_C
	q.Option_D = src.Option_D
	q.Option_E = src.Option_E
	q.Answer = src.Answer
}
