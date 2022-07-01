package router

import (
	QuestionRouter "github.com/isazobu/dailyQuestionsAPI/questions/router"
	"github.com/labstack/echo/v4"
)

func RegisterAllRoute(v1 *echo.Group) {
	QuestionRouter.QuestionRegister(v1)

}
