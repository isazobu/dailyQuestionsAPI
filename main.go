package main

import (
	db "github.com/isazobu/dailyQuestionsAPI/database"
	questionController "github.com/isazobu/dailyQuestionsAPI/questions/controller"
	questionRepository "github.com/isazobu/dailyQuestionsAPI/questions/repository/mongodb"
	questionService "github.com/isazobu/dailyQuestionsAPI/questions/services"
	"github.com/isazobu/dailyQuestionsAPI/router"
	setup "github.com/isazobu/dailyQuestionsAPI/setup"
)

func main() {

	r := setup.New()
	client := db.ConnectDB()

	qr := questionRepository.NewQuestionRepository(db.GetCollection(client, "question"))
	qs := questionService.NewQuestionService(qr)
	qc := questionController.NewQuestionController(qs)
	// Start server

	router.InitRoute(r)
	r.Logger.Fatal(r.Start(":3000"))
}
