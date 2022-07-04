package questionrepo

import (
	"fmt"
	"net/url"

	dto "github.com/isazobu/dailyQuestionsAPI/questions/dtos"
	"github.com/isazobu/dailyQuestionsAPI/questions/models"
	"gopkg.in/mgo.v2/bson"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	AddQuestion(question dto.CreateQuestion) (*mongo.InsertOneResult, error)
	GetQuestionsByFilter(params url.Values) ([]models.Question, error)
	GetQuestionById(id string) (models.Question, error)
	GetQuestions() ([]models.Question, error)
	UpdateQuestion(question models.Question) (*mongo.UpdateResult, error)
	DeleteQuestion(id string) error
}

type questionRepo struct {
	collection *mongo.Collection
}

func NewQuestionRepository(col *mongo.Collection) Repo {
	return &questionRepo{
		collection: col,
	}
}

func (q questionRepo) AddQuestion(question dto.CreateQuestion) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := q.collection.InsertOne(ctx, question)
	return res, err
}

func (q questionRepo) GetQuestionsByFilter(params url.Values) ([]models.Question, error) {
	filter := make(bson.M)
	for key, value := range params {
		filter[key] = value
	}
	return q.GetQuestionsBySpecifiedFilter(filter)
}

func (q questionRepo) GetQuestionById(id string) (models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": bson.ObjectIdHex(id)}
	var question models.Question
	err := q.collection.FindOne(ctx, filter).Decode(&question)
	return question, err
}

func (q questionRepo) GetQuestions() ([]models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := q.collection.Find(ctx, bson.M{})
	defer cur.Close(ctx)

	if err != nil {
		return make([]models.Question, 0), err
	}

	var questions []models.Question
	if err = cur.All(ctx, &questions); err != nil {
		return make([]models.Question, 0), err
	}
	if questions == nil {
		return make([]models.Question, 0), nil
	}
	return questions, nil
}

func (q questionRepo) UpdateQuestion(question models.Question) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": bson.ObjectIdHex(question.Id.Hex())}
	update := bson.M{"$set": question}
	res, err := q.collection.UpdateOne(ctx, filter, update)
	return res, err
}

func (q questionRepo) DeleteQuestion(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": bson.ObjectIdHex(id)}
	_, err := q.collection.DeleteOne(ctx, filter)
	return err
}

func (q questionRepo) GetQuestionsBySpecifiedFilter(filter bson.M) ([]models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println(filter)
	cur, err := q.collection.Find(ctx, filter)
	if err != nil {
		return make([]models.Question, 0), err
	}
	defer cur.Close(ctx)

	var questions []models.Question
	if err = cur.All(ctx, &questions); err != nil {
		return make([]models.Question, 0), err
	}
	if questions == nil {
		return make([]models.Question, 0), nil
	}
	return questions, nil
}
