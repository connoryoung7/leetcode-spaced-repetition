package services

import (
	"leetcode-spaced-repetition/models"
	"leetcode-spaced-repetition/repositories"
	"time"

	"github.com/google/uuid"
)

type QuestionService struct {
	questionRepo repositories.QuestionRepository
}

func NewQuestionsService(questionsRepo repositories.QuestionRepository) *QuestionService {
	return &QuestionService{
		questionRepo: questionsRepo,
	}
}

func (s QuestionService) GetQuestionByID(ID int) (*models.Question, error) {
	return s.questionRepo.GetQuestionByID(ID)
}

func (s QuestionService) GetQuestionSubmissions() ([]models.QuestionSubmission, error) {
	return s.questionRepo.GetQuestionSubmissions(1)
}

func (s QuestionService) GetAllQuestionTags() ([]string, error) {
	return s.questionRepo.GetAllQuestionTags()
}

func (s QuestionService) GetTagsForQuestion(ID int) ([]string, error) {
	return s.questionRepo.GetTagsForQuestion(ID)
}

func (s QuestionService) GetAllQuestionsPastReviewDate(limit uint) ([]models.Question, error) {
	return s.questionRepo.GetAllQuestionsPastReviewDate(limit)
}

func (s QuestionService) GetAllSubmissionsForQuestion(questionID int) ([]models.QuestionSubmission, error) {
	submissions, err := s.questionRepo.GetQuestionSubmissions(questionID)
	if err != nil {
		return []models.QuestionSubmission{}, err
	}

	return submissions, nil
}

func (s QuestionService) SaveQuestionSubmission(questionID int, userID uuid.UUID, date time.Time, timeTaken time.Duration, confidenceLevel models.ConfidenceLevel) error {
	return s.questionRepo.SaveQuestionSubmission(questionID, userID, date, timeTaken, confidenceLevel)
}
