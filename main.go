package main

import (
	db "github.com/isazobu/dailyQuestionsAPI/database"
	"github.com/isazobu/dailyQuestionsAPI/questions/models"
	questionRepository "github.com/isazobu/dailyQuestionsAPI/questions/repository/mongodb"
	QuestionRouter "github.com/isazobu/dailyQuestionsAPI/questions/router"
	questionService "github.com/isazobu/dailyQuestionsAPI/questions/services"
	setup "github.com/isazobu/dailyQuestionsAPI/setup"
)

func main() {

	r := setup.New()
	QuestionRouter.QuestionRegister(r)
	client := db.ConnectDB()

	qr := questionRepository.NewQuestionRepository(db.GetCollection(client, "question"))
	qs := questionService.NewQuestionService(qr)
	qs.AddQuestion(models.Question{})
	// Start server
	r.Logger.Fatal(r.Start(":3000"))
}
