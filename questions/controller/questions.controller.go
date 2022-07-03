package QuestionController

import (
	"net/http"

	"github.com/isazobu/dailyQuestionsAPI/questions/models"
	qs "github.com/isazobu/dailyQuestionsAPI/questions/services"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	AddQuestion(ctx echo.Context) error
}

type questionController struct {
	qs qs.QuestionService
}

func NewQuestionController(qs qs.QuestionService) Controller {

	return &questionController{qs: qs}
}

func (q questionController) AddQuestion(ctx echo.Context) error {
	var newQuestion models.Question
	if error := ctx.Bind(&newQuestion); error != nil {
		return ctx.JSON(http.StatusNotAcceptable, nil)
	}
	if _, error := q.qs.AddQuestion(newQuestion); error != nil {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}

	return ctx.JSON(http.StatusCreated, newQuestion)
}
