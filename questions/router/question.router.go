package QuestionRouter

import (
	QuestionController "github.com/isazobu/dailyQuestionsAPI/questions/controller"
	"github.com/labstack/echo/v4"
)

func QuestionRegister(g *echo.Group, q QuestionController.Controller) {
	g.POST("/questions", q.AddQuestion)
	g.GET("/questions", q.GetQuestions)
	g.GET("/questions/:id", q.GetQuestionById)
}
