package QuestionController

import (
	"fmt"
	"net/http"
	"time"

	dto "github.com/isazobu/dailyQuestionsAPI/questions/dtos"
	qs "github.com/isazobu/dailyQuestionsAPI/questions/services"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller interface {
	AddQuestion(ctx echo.Context) error
	GetQuestions(ctx echo.Context) error
	GetQuestionById(ctx echo.Context) error
	UpdateQuestion(ctx echo.Context) error
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
	if err := ctx.Bind(&newQuestion); err != nil {
		return ctx.JSON(http.StatusNotAcceptable, err.Error())
	}
	fmt.Printf("%+v\n", newQuestion)
	if _, err := q.qs.AddQuestion(newQuestion); err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	return ctx.JSON(http.StatusCreated, newQuestion)
}

func (q questionController) GetQuestions(ctx echo.Context) error {
	if len(ctx.QueryParams()) == 0 {
		questions, err := q.qs.GetQuestions()
		if err != nil {
			return ctx.JSON(http.StatusNotAcceptable, err.Error())
		}
		return ctx.JSON(http.StatusOK, questions)
	} else {
		questions, err := q.qs.GetQuestionsByFilter(ctx.QueryParams())
		if err != nil {
			return ctx.JSON(http.StatusNotAcceptable, err.Error())
		}
		return ctx.JSON(http.StatusOK, questions)
	}
}

func (q questionController) GetQuestionById(ctx echo.Context) error {
	fmt.Println(ctx.Param("id"))
	question, err := q.qs.GetQuestionById(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, question)
}

func (q questionController) UpdateQuestion(ctx echo.Context) error {
	updateQuestion := dto.UpdateQuestion{}
	updateQuestion.UpdatedAt = time.Now()
	if err := ctx.Bind(&updateQuestion); err != nil {
		return ctx.JSON(http.StatusNotAcceptable, err.Error())
	}
	objID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNotAcceptable, err.Error())
	}
	updateQuestion.Id = objID
	if _, err := q.qs.UpdateQuestion(updateQuestion); err != nil {
		return ctx.JSON(http.StatusNotModified, err.Error())
	}
	return ctx.JSON(http.StatusOK, updateQuestion)
}
