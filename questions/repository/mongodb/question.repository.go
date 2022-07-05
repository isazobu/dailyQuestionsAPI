package questionrepo

import (
	"fmt"
	"net/url"
	"strings"

	dto "github.com/isazobu/dailyQuestionsAPI/questions/dtos"
	"github.com/isazobu/dailyQuestionsAPI/questions/models"
	"gopkg.in/mgo.v2/bson"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	AddQuestion(question dto.CreateQuestion) (*mongo.InsertOneResult, error)
	GetQuestionsByFilter(params url.Values) ([]models.Question, error)
	GetQuestionById(id string) (models.Question, error)
	GetQuestions() ([]models.Question, error)
	UpdateQuestion(question dto.UpdateQuestion) (*mongo.UpdateResult, error)
	DeleteQuestion(id string) (*mongo.DeleteResult, error)
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
	criterias := []bson.M{}
	for key, value := range params {
		operands := make(bson.M)
		if len(value) == 1 {
			value = strings.Split(value[0], ",")
		}
		x := []bson.M{}
		for _, val := range value {
			y := make(bson.M)
			y[key] = val
			x = append(x, y)
		}
		operands["$or"] = x
		criterias = append(criterias, operands)
	}
	filter["$and"] = criterias

	//basic or operation for mongodb
	//filter := bson.M{"$or": []bson.M{{"category": "Matematik"}, {"category": "Türkçe"}}}

	return q.GetQuestionsBySpecifiedFilter(filter)
}

func (q questionRepo) GetQuestionById(id string) (models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var question models.Question
	obj_id, err_id := primitive.ObjectIDFromHex(id)
	if err_id != nil {
		return question, err_id
	}
	filter := bson.M{"_id": obj_id}
	err := q.collection.FindOne(ctx, filter).Decode(&question)
	return question, err
}

func (q questionRepo) GetQuestions() ([]models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := q.collection.Find(ctx, bson.M{})

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

func (q questionRepo) UpdateQuestion(question dto.UpdateQuestion) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	obj_id, err1 := primitive.ObjectIDFromHex(question.Id.Hex())
	if err1 != nil {
		return nil, err1
	}
	filter := bson.M{"_id": obj_id}
	update := bson.M{"$set": question}
	res, err := q.collection.UpdateOne(ctx, filter, update)
	return res, err
}

func (q questionRepo) DeleteQuestion(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	obj_id, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		return nil, err1
	}
	filter := bson.M{"_id": obj_id}
	res, err := q.collection.DeleteOne(ctx, filter)
	return res, err
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
