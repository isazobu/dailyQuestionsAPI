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

func QuestionRegister(e *echo.Echo) {

	e.GET("/questions", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

}
