package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"leetcode-spaced-repetition/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

type QuestionPostgresRepository struct {
	db *sql.DB
}

// SaveQuestionSubmission implements QuestionRepository.
func (r QuestionPostgresRepository) SaveQuestionSubmission(c context.Context, questionID int, userID uuid.UUID, date time.Time, timeTaken time.Duration, confidenceLevel models.ConfidenceLevel) error {
	_, err := r.db.ExecContext(
		c,
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

func (r QuestionPostgresRepository) GetQuestions(ctx context.Context, tags []string, page int, limit int) ([]models.Question, error) {
	var questions []models.Question

	rows, err := r.db.QueryContext(
		ctx, `SELECT id, title, slug, difficulty FROM questions WHERE id IN (
		SELECT questionId FROM questionTags WHERE tag IN ($1)
	) ORDER BY id LIMIT $2`, strings.Join(tags, ","), limit)
	if err != nil {
		return questions, err
	}
	defer rows.Close()

	for rows.Next() {
		var question models.Question
		err = rows.Scan(&question.ID, &question.Title, &question.Slug, &question.Difficulty)
		if err != nil {
			return questions, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func (r QuestionPostgresRepository) GetQuestionByID(ctx context.Context, ID int) (*models.Question, error) {
	var id int
	var title string
	var slug string
	var difficulty int

	row := r.db.QueryRowContext(ctx, "SELECT id, title, slug, difficulty FROM questions WHERE id = $1", ID)
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

func (r QuestionPostgresRepository) GetQuestionStatsByID(ctx context.Context, ID int) (*models.QuestionSubmissionUserStats, error) {
	return nil, nil
}

func (r QuestionPostgresRepository) GetQuestionSubmissions(ctx context.Context, questionID int) ([]models.QuestionSubmission, error) {
	rows, err := r.db.QueryContext(
		ctx,
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

func (r QuestionPostgresRepository) GetAllQuestionTags(ctx context.Context) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT DISTINCT(tag) FROM questionTags ORDER BY tag")
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

func (r QuestionPostgresRepository) GetTagsForQuestion(ctx context.Context, ID int) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT tag FROM questionTags WHERE questionId = $1", ID)
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
