package QuestionRouter

import (
	"fmt"
	"net/http"

	"github.com/isazobu/dailyQuestionsAPI/questions/models"
	repository "github.com/isazobu/dailyQuestionsAPI/questions/repository/mongodb"
	"github.com/labstack/echo/v4"
)

func insertQuestion(c echo.Context) error {
	u := &models.Question{}
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println(u)

	question, err := repository.Insert(models.Question{Title: u.Title, Image: u.Image})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(question)
	return c.JSON(http.StatusOK, question)
}

func QuestionRegister(v1 *echo.Group) {
	fmt.Println("Girdim")
	questions := v1.Group("/questions")
	// questions.GET("/", GetQuestions)
	questions.POST("/", insertQuestion)
	// questions.GET("/:id", GetQuestion)
	// questions.PUT("/:id", UpdateQuestion)
	// questions.DELETE("/:id", DeleteQuestion)

}
