package questionrepo

import (
	"github.com/isazobu/dailyQuestionsAPI/questions/models"
	"gopkg.in/mgo.v2/bson"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	AddQuestion(question models.Question) (*mongo.InsertOneResult, error)
	GetQuestionsByCategory(category string) ([]models.Question, error)
	GetQuestionsByDiffuculty(diffuculty string) ([]models.Question, error)
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

// var questionCollection *mongo.Collection = db.GetCollection(db.DB, "question")

func (q questionRepo) AddQuestion(question models.Question) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := q.collection.InsertOne(ctx, question)
	return res, err
}

func (q questionRepo) GetQuestionsByCategory(category string) ([]models.Question, error) {
	filter := bson.M{"category": category}
	return q.GetQuestionsBySpecifiedFilter(filter)
}
func (q questionRepo) GetQuestionsByDiffuculty(diffuculty string) ([]models.Question, error) {
	filter := bson.M{"diffuculty": diffuculty}
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

func (q questionRepo) UpdateQuestion(question models.Question) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": bson.ObjectIdHex(question.Id)}
	update := bson.M{"$set": question}
	res, err := q.collection.UpdateOne(ctx, filter, update)
	return res, err
}

func (q questionRepo) DeleteQuestion(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": bson.ObjectIdHex(id)}
	res, err := q.collection.DeleteOne(ctx, filter)
	return res, err
}

func (q questionRepo) GetQuestionsBySpecifiedFilter(filter bson.M) ([]models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
