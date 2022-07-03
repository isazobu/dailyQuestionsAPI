package QuestionRouter

import (
	"net/http"

	QuestionController "github.com/isazobu/dailyQuestionsAPI/questions/controller"
	"github.com/labstack/echo/v4"
)

func QuestionRegister(g *echo.Group, q QuestionController.Controller) {
	g.POST("/questions", q.AddQuestion)
	g.GET("/questions", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.QueryParam("id"))
	})
}
