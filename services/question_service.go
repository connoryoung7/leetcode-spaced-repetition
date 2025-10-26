package services

import (
	"context"
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

func (s QuestionService) GetQuestions(ctx context.Context, tags []string, page int, limit int) ([]models.Question, error) {
	return s.questionRepo.GetQuestions(ctx, tags, page, limit)
}

func (s QuestionService) GetQuestionByID(c context.Context, ID int) (*models.Question, error) {
	return s.questionRepo.GetQuestionByID(c, ID)
}

func (s QuestionService) GetQuestionSubmissions(c context.Context, questionId int) ([]models.QuestionSubmission, error) {
	return s.questionRepo.GetQuestionSubmissions(c, questionId)
}

func (s QuestionService) GetAllQuestionTags(c context.Context) ([]string, error) {
	return s.questionRepo.GetAllQuestionTags(c)
}

func (s QuestionService) GetTagsForQuestion(c context.Context, ID int) ([]string, error) {
	return s.questionRepo.GetTagsForQuestion(c, ID)
}

func (s QuestionService) GetAllQuestionsPastReviewDate(c context.Context, limit uint) ([]models.Question, error) {
	return s.questionRepo.GetAllQuestionsPastReviewDate(c, limit)
}

func (s QuestionService) GetAllSubmissionsForQuestion(c context.Context, questionID int) ([]models.QuestionSubmission, error) {
	submissions, err := s.questionRepo.GetQuestionSubmissions(c, questionID)
	if err != nil {
		return []models.QuestionSubmission{}, err
	}

	return submissions, nil
}

func (s QuestionService) SaveQuestionSubmission(
	c context.Context,
	questionID int,
	userID uuid.UUID,
	date time.Time,
	timeTaken time.Duration,
	confidenceLevel models.ConfidenceLevel,
) error {
	return s.questionRepo.SaveQuestionSubmission(c, questionID, userID, date, timeTaken, confidenceLevel)
}
