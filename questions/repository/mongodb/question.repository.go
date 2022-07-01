package repository

import (
	db "github.com/isazobu/dailyQuestionsAPI/database"
	"github.com/isazobu/dailyQuestionsAPI/questions/models"

	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var questionCollection *mongo.Collection = db.GetCollection(db.DB, "question")

func Insert(question models.Question) (*mongo.InsertOneResult, error) {
	fmt.Println("Starting the Insert process")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := questionCollection.InsertOne(ctx, question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Insert successful")
	return result, err
}
