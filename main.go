package main

import (
	"fmt"
	"net/http"
	"strconv"

	db "github.com/isazobu/dailyQuestionsAPI/database"
	"github.com/isazobu/dailyQuestionsAPI/questions/models"
	repository "github.com/isazobu/dailyQuestionsAPI/questions/repository/mongodb"
	router "github.com/isazobu/dailyQuestionsAPI/router"
	setup "github.com/isazobu/dailyQuestionsAPI/setup"
	"github.com/labstack/echo/v4"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {

	r := setup.New()
	v1 := r.Group("/api")
	router.RegisterAllRoute(v1)

	db.ConnectDB()

	// Start server
	r.Logger.Fatal(r.Start(":3000"))
}
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
