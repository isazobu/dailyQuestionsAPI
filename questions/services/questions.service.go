package questionservice

import (
	models "github.com/isazobu/dailyQuestionsAPI/questions/models"
	questionrepo "github.com/isazobu/dailyQuestionsAPI/questions/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionService interface {
	AddQuestion(question models.Question) (*mongo.InsertOneResult, error)
	GetQuestionsByCategory(category string) ([]models.Question, error)
	GetQuestionById(id string) (models.Question, error)
	UpdateQuestion(question models.Question) (*mongo.UpdateResult, error)
	DeleteQuestion(id string) error
}
type questionService struct {
	Repo questionrepo.Repo
}

func NewQuestionService(qRepo questionrepo.Repo) QuestionService {
	return &questionService{Repo: qRepo}
}

func (q questionService) AddQuestion(question models.Question) (*mongo.InsertOneResult, error) {
	return q.Repo.AddQuestion(question)
}

func (q questionService) GetQuestionsByCategory(category string) ([]models.Question, error) {
	return q.Repo.GetQuestionsByCategory(category)
}

func (q questionService) GetQuestionById(id string) (models.Question, error) {
	return q.Repo.GetQuestionById(id)
}

func (q questionService) UpdateQuestion(question models.Question) (*mongo.UpdateResult, error) {
	return q.Repo.UpdateQuestion(question)
}

func (q questionService) DeleteQuestion(id string) error {
	return q.Repo.DeleteQuestion(id)
}
