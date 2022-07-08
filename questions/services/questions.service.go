package questionservice

import (
	"net/url"

	dto "github.com/isazobu/dailyQuestionsAPI/questions/dtos"
	models "github.com/isazobu/dailyQuestionsAPI/questions/models"
	questionrepo "github.com/isazobu/dailyQuestionsAPI/questions/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionService interface {
	AddQuestion(question dto.QuestionDTO) (*mongo.InsertOneResult, error)
	GetQuestions() ([]models.Question, error)
	GetQuestionsByFilter(params url.Values) ([]dto.QuestionDTO, error)
	GetQuestionById(id string) (dto.QuestionDTO, error)
	UpdateQuestion(question dto.QuestionDTO, id string) (*mongo.UpdateResult, error)
	DeleteQuestion(id string) (*mongo.DeleteResult, error)
}
type questionService struct {
	Repo questionrepo.Repo
}

func NewQuestionService(qRepo questionrepo.Repo) QuestionService {
	return &questionService{Repo: qRepo}
}

func (q questionService) AddQuestion(question dto.QuestionDTO) (*mongo.InsertOneResult, error) {
	newQuestion := question.MapToQuestionModel()
	newQuestion.Id = primitive.NilObjectID
	return q.Repo.AddQuestion(newQuestion)
}

func (q questionService) GetQuestions() ([]models.Question, error) {
	return q.Repo.GetQuestions()
}

func (q questionService) GetQuestionsByFilter(params url.Values) ([]dto.QuestionDTO, error) {
	res, err := q.Repo.GetQuestionsByFilter(params)
	if err != nil {
		return []dto.QuestionDTO{}, err
	}
	resDTO := make([]dto.QuestionDTO, 0)
	for _, val := range res {
		var x dto.QuestionDTO
		(&x).MapFromQuestionModel(val)
		resDTO = append(resDTO, x)
	}
	return resDTO, nil
}

func (q questionService) GetQuestionById(id string) (dto.QuestionDTO, error) {
	res, err := q.Repo.GetQuestionById(id)
	if err != nil {
		return dto.QuestionDTO{}, err
	}
	var retval dto.QuestionDTO
	(&retval).MapFromQuestionModel(res)
	return retval, nil
}

func (q questionService) UpdateQuestion(question dto.QuestionDTO, id string) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updated_question := question.MapToQuestionModel()
	updated_question.Id = objID
	return q.Repo.UpdateQuestion(updated_question)
}

func (q questionService) DeleteQuestion(id string) (*mongo.DeleteResult, error) {
	return q.Repo.DeleteQuestion(id)
}
