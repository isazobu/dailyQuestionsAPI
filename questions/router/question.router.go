package QuestionRouter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func QuestionRegister(e *echo.Echo) {

	e.GET("/questions", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

}
