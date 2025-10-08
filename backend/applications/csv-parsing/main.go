package main

import (
	"encoding/csv"
	"fmt"
	"leetcode-spaced-repetition/internal"
	"leetcode-spaced-repetition/models"
	"leetcode-spaced-repetition/repositories"
	"leetcode-spaced-repetition/services"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	ProblemNumberIdx int = iota
	DateIdx
	TimeTakenIdx
	ConfidenceLevelIdx
)

const dateFormat string = "2006-01-02"

type recordSubmission struct {
	questionNumber  int
	timeTaken       time.Duration
	submissionDate  time.Time
	confidenceLevel models.ConfidenceLevel
}

func main() {
	config, err := internal.GetConfig()
	if err != nil {
		fmt.Println("Cannot access configuration")
		return
	}
	db, err := internal.GetDBConnFromConfig(config)
	if err != nil {
		fmt.Println("Cannot construct DB")
		return
	}

	fmt.Print("Constructing different domain layers...")

	questionsRepo := repositories.NewQuestionPostgresRepository(db)
	questionsService := services.NewQuestionsService(questionsRepo)

	file, err := os.Open("leetcode_submissions.csv")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	fmt.Println("Reading the leetcode submission files")

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Cannot read csv file")
		panic(err)
	}

	fmt.Printf("There are %d rows in the file\n", len(records))

	for i := 1; i < len(records); i++ {
		fmt.Printf("Record: %+v\n", records[i])
		submission, err := validateQuestionSubmission(records[i])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		userID, _ := uuid.NewUUID()

		err = questionsService.SaveQuestionSubmission(submission.questionNumber, userID, submission.submissionDate, submission.timeTaken, submission.confidenceLevel)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Successfully created question submission")
		}
	}
}

func validateQuestionSubmission(r []string) (recordSubmission, error) {
	problemNum := r[ProblemNumberIdx]
	date := r[DateIdx]
	timeTaken := r[TimeTakenIdx]
	confidenceLevelStr := r[ConfidenceLevelIdx]

	questionNum, err := strconv.ParseInt(problemNum, 10, 0)
	if err != nil {
		return recordSubmission{}, fmt.Errorf("'%s' is not a valid question number", problemNum)
	}

	timeTakenDuration, err := time.ParseDuration(strings.ReplaceAll(timeTaken, " ", ""))
	if err != nil {
		return recordSubmission{}, fmt.Errorf("'%s' is not a valid time duration", timeTaken)
	}

	dateTime, err := time.Parse(dateFormat, date)
	if err != nil {
		return recordSubmission{}, fmt.Errorf("'%s' is not a valid date format", date)
	}

	confidenceLevel, err := models.DetermineConfidenceLevelFromString(confidenceLevelStr)
	if err != nil {
		return recordSubmission{}, fmt.Errorf("'%s' is not a valid confidence level", confidenceLevelStr)
	}

	return recordSubmission{
		questionNumber:  int(questionNum),
		timeTaken:       timeTakenDuration,
		submissionDate:  dateTime,
		confidenceLevel: confidenceLevel,
	}, nil
}
