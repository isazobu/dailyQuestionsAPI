package QuestionController

import (
	"fmt"
	"net/http"
	"time"

	dto "github.com/isazobu/dailyQuestionsAPI/questions/dtos"
	qs "github.com/isazobu/dailyQuestionsAPI/questions/services"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	AddQuestion(ctx echo.Context) error
	GetQuestions(ctx echo.Context) error
}

type questionController struct {
	qs qs.QuestionService
}

func NewQuestionController(qs qs.QuestionService) Controller {

	return &questionController{qs: qs}
}

func (q questionController) AddQuestion(ctx echo.Context) error {
	var newQuestion dto.CreateQuestion
	newQuestion.CreatedAt = time.Now()
	if error := ctx.Bind(&newQuestion); error != nil {
		return ctx.JSON(http.StatusNotAcceptable, nil)
	}
	fmt.Printf("%+v\n", newQuestion)
	if _, error := q.qs.AddQuestion(newQuestion); error != nil {
		return ctx.JSON(http.StatusUnauthorized, nil)
	}

	return ctx.JSON(http.StatusCreated, newQuestion)
}

func (q questionController) GetQuestions(ctx echo.Context) error {
	if len(ctx.QueryParams()) == 0 {
		questions, error := q.qs.GetQuestions()
		if error != nil {
			return ctx.JSON(http.StatusNotAcceptable, nil)
		}
		return ctx.JSON(http.StatusOK, questions)
	} else {
		questions, error := q.qs.GetQuestionsByFilter(ctx.QueryParams())
		if error != nil {
			return ctx.JSON(http.StatusNotAcceptable, nil)
		}
		return ctx.JSON(http.StatusOK, questions)
	}
}
