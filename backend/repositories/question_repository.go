package repositories

import (
	"leetcode-spaced-repetition/models"
	"time"

	"github.com/google/uuid"
)

type QuestionRepository interface {
	GetQuestionSubmissions(questionID int) ([]models.QuestionSubmission, error)
	GetQuestionByID(id int) (*models.Question, error)
	GetQuestionStatsByID(id int) (*models.QuestionSubmissionUserStats, error)
	GetAllQuestionsPastReviewDate(limit uint) ([]models.Question, error)
	GetQuestions() ([]models.Question, error)
	GetAllQuestionTags() ([]string, error)
	GetTagsForQuestion(ID int) ([]string, error)
	SaveQuestion(question models.Question) error
	SaveQuestionTag(questionId int, tag string) error
	SaveQuestionSubmission(questionID int, userID uuid.UUID, date time.Time, timeTaken time.Duration, confidenceLevel models.ConfidenceLevel) error
}
