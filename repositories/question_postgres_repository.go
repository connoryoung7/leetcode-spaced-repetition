package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"leetcode-spaced-repetition/models"
	"time"

	"github.com/google/uuid"
)

type QuestionPostgresRepository struct {
	db *sql.DB
}

// SaveQuestionSubmission implements QuestionRepository.
func (r QuestionPostgresRepository) SaveQuestionSubmission(c context.Context, questionID int, userID uuid.UUID, date time.Time, timeTaken time.Duration, confidenceLevel models.ConfidenceLevel) error {
	fmt.Printf("timeTaken = %d\n", int(timeTaken.Seconds()))
	_, err := r.db.Exec(
		`INSERT INTO questionSubmissions (questionId, userId, submissionDate, timeTaken, confidenceLevel) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (questionId, userId, submissionDate) DO NOTHING`,
		questionID,
		userID,
		date,
		fmt.Sprintf("%d seconds", int64(timeTaken.Seconds())),
		confidenceLevel,
	)

	return err
}

// GetAllQuestionsPastReviewDate implements QuestionRepository.
func (r QuestionPostgresRepository) GetAllQuestionsPastReviewDate(c context.Context, limit uint) ([]models.Question, error) {
	var questions []models.Question
	return questions, nil
}

func NewQuestionPostgresRepository(db *sql.DB) *QuestionPostgresRepository {
	return &QuestionPostgresRepository{
		db: db,
	}
}

func (r QuestionPostgresRepository) GetQuestions(c context.Context) ([]models.Question, error) {
	return []models.Question{}, nil
}

func (r QuestionPostgresRepository) GetQuestionByID(c context.Context, ID int) (*models.Question, error) {
	var id int
	var title string
	var slug string
	var difficulty int

	row := r.db.QueryRow("SELECT id, title, slug, difficulty FROM questions WHERE id = $1", ID)
	switch err := row.Scan(&id, &title, &slug, &difficulty); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &models.Question{
			ID:         id,
			Title:      title,
			Slug:       slug,
			Difficulty: models.QuestionDifficulty(difficulty),
		}, nil
	default:
		return nil, err
	}
}

func (r QuestionPostgresRepository) GetQuestionStatsByID(c context.Context, ID int) (*models.QuestionSubmissionUserStats, error) {
	return nil, nil
}

func (r QuestionPostgresRepository) GetQuestionSubmissions(c context.Context, questionID int) ([]models.QuestionSubmission, error) {
	rows, err := r.db.Query(
		"SELECT id, questionID, submissionDate, EXTRACT(EPOCH  FROM timeTaken), confidenceLevel FROM questionSubmissions WHERE questionId = $1 ORDER BY submissionDate DESC",
		questionID,
	)
	if err != nil {
		return []models.QuestionSubmission{}, err
	}
	defer rows.Close()

	var submissions []models.QuestionSubmission
	for rows.Next() {
		var sub models.QuestionSubmission

		if err := rows.Scan(&sub.ID, &sub.QuestionID, &sub.Date, &sub.TimeTaken, &sub.ConfidenceLevel); err != nil {
			return []models.QuestionSubmission{}, err
		}

		submissions = append(submissions, sub)
	}

	return submissions, nil
}

func (r QuestionPostgresRepository) SaveQuestion(c context.Context, q models.Question) error {
	_, err := r.db.Exec(
		"INSERT INTO questions (id, title, slug, description, difficulty) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO NOTHING",
		q.ID, q.Title, q.Slug, q.Description, q.Difficulty,
	)

	return err
}

func (r QuestionPostgresRepository) SaveQuestionTag(c context.Context, questionId int, tag string) error {
	_, err := r.db.Exec(
		"INSERT INTO questionTags (questionId, tag) VALUES ($1, $2) ON CONFLICT (questionId, tag) DO NOTHING",
		questionId, tag,
	)

	return err
}

func (r QuestionPostgresRepository) GetAllQuestionTags(c context.Context) ([]string, error) {
	rows, err := r.db.Query("SELECT DISTINCT(tag) FROM questionTags ORDER BY tag")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var tags []string

	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return []string{}, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r QuestionPostgresRepository) GetTagsForQuestion(c context.Context, ID int) ([]string, error) {
	rows, err := r.db.Query("SELECT tag FROM questionTags WHERE questionId = $1", ID)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var tags []string

	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return []string{}, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
