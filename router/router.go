package router

import (
	QuestionController "github.com/isazobu/dailyQuestionsAPI/questions/controller"
	QuestionRouter "github.com/isazobu/dailyQuestionsAPI/questions/router"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, q QuestionController.Controller) {
	g := e.Group("/api")
	QuestionRouter.QuestionRegister(g, q)
}
