package QuestionController

import (
	"github.com/isazobu/dailyQuestionsAPI/questions/models"
)

type QuestionController interface {
	AddQuestion(question models.Question)
}

type Handler struct {
	userStore    user.Store
	articleStore article.Store
}

func NewHandler(us user.Store, as article.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
	}
}
