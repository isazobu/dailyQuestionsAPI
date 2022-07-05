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
	DeleteQuestion(ctx echo.Context) error
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
	res, err := q.qs.AddQuestion(newQuestion)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return ctx.JSON(http.StatusCreated, "Question is inserted but its id cannot be converted to primitive.ObjectID")
	}
	question, err := q.qs.GetQuestionById(oid.Hex())
	if err != nil {
		return ctx.JSON(http.StatusCreated, "Question is inserted but an error occur while fetching it from db: "+err.Error())
	}
	return ctx.JSON(http.StatusCreated, question)
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

func (q questionController) DeleteQuestion(ctx echo.Context) error {
	if res, err := q.qs.DeleteQuestion(ctx.Param("id")); err != nil || res.DeletedCount == 0 {
		return ctx.JSON(http.StatusNoContent, nil)
	}
	return ctx.JSON(http.StatusOK, "Successfully deleted.")
}
