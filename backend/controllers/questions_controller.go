package controllers

import (
	"leetcode-spaced-repetition/models"
	"leetcode-spaced-repetition/services"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const dateRegexString = `^\d{4}-\d{2}-\d{2}T\d{2}`

type leetCodeQuestionRequest struct {
	ID string `uri:"id" binding:"required,number"`
}

type saveQuestionSubmissionRequest struct {
	QuestionID      int `json:"questionId binding:"required,number"`
	TimeTaken       int `json:"timeTaken" binding:"required,number"`
	ConfidenceLevel int `json:"confidenceLevel"`
}

var validDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(string)

	if !ok {
		return false
	}
	match, _ := regexp.MatchString(dateRegexString, date)
	return match
}

type QuestionsController struct {
	questionsService services.QuestionService
}

func RegisterRoutes(r *gin.Engine, questionsService *services.QuestionService) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("date", validDate)
	}

	questionsController := QuestionsController{questionsService: *questionsService}

	questionsGroup := r.Group("/questions")
	questionsGroup.GET("tags", questionsController.GetAllQuestionTags)
	questionsGroup.GET(":id", questionsController.GetQuestionByID)

	individualQuestionsGroup := questionsGroup.Group(":id")
	individualQuestionsGroup.GET("submissions", questionsController.getQuestionSubmissionsByID)
}

func (c QuestionsController) GetQuestionByID(context *gin.Context) {
	var request leetCodeQuestionRequest
	if err := context.ShouldBindUri(&request); err != nil {
		context.JSON(400, gin.H{
			"error": "The id of the question must be a valid integer.",
		})
		return
	}

	intId, err := strconv.Atoi(request.ID)
	if err != nil {
		context.JSON(400, gin.H{
			"error": "The id of the question must be a valid integer.",
		})
		return
	}

	question, err := c.questionsService.GetQuestionByID(intId)
	if err != nil {
		context.JSON(500, gin.H{
			"error": "An internal server has occurred.",
		})
		return
	}
	if question == nil {
		context.JSON(404, gin.H{
			"message": "No question is associated with this code.",
		})
		return
	}

	tags, err := c.questionsService.GetTagsForQuestion(intId)
	if err != nil {
		context.JSON(500, gin.H{
			"error": "An internal server error has occurred.",
		})
		return
	}

	question.Tags = tags

	context.JSON(200, *question)
}

func (c QuestionsController) getQuestionSubmissionsByID(context *gin.Context) {
	var request leetCodeQuestionRequest

	if err := context.ShouldBindUri(&request); err != nil {
		context.JSON(400, gin.H{
			"error": "The id of the question must be a valid integer.",
		})
	}

	questionId, _ := strconv.ParseInt(request.ID, 10, 0)

	var resp models.Pagaination[models.QuestionSubmission]

	result, err := c.questionsService.GetAllSubmissionsForQuestion(int(questionId))
	if err != nil {
		context.JSON(500, gin.H{
			"error": "An internal server error occurred.",
		})
	}

	resp.Data = result

	context.JSON(200, resp)
}

func (c QuestionsController) GetAllQuestionTags(context *gin.Context) {
	tags, err := c.questionsService.GetAllQuestionTags()
	if err != nil {
		context.JSON(500, gin.H{
			"error": "An internal server error has occurred.",
		})
	}

	context.JSON(200, gin.H{
		"tags": tags,
	})
}

func (c QuestionsController) SaveQuestionController(context *gin.Context) {
	var questionSubmissionRequest saveQuestionSubmissionRequest
	if err := context.ShouldBindBodyWithJSON(&questionSubmissionRequest); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid request body",
		})
	}

	if err = c.questionsService.SaveQuestionSubmission(
		questionSubmissionRequest.QuestionID,
		uuid.Must(uuid.New()),
		
	)

	context.JSON(200, gin.H{
		"message": "Successfully saved question submission",
	})
}
