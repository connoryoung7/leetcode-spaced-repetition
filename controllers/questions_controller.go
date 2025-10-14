package controllers

import (
	"leetcode-spaced-repetition/models"
	"leetcode-spaced-repetition/services"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const dateRegexString = `^\d{4}-\d{2}-\d{2}T\d{2}`

type leetCodeQuestionRequest struct {
	ID string `uri:"id" binding:"required,number" validate:"gte=1"`
}

type saveQuestionSubmissionRequest struct {
	QuestionID      int `json:"questionId" binding:"required,number" validate:"gte=1"`
	TimeTaken       int `json:"timeTaken" binding:"required,number" validate:"gte=0"`
	ConfidenceLevel int `json:"confidenceLevel" binding:"required,number" validate:"gte=1,lte=5"`
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
	questionsGroup.POST("submissions", questionsController.SaveQuestionSubmission)
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

	question, err := c.questionsService.GetQuestionByID(context, intId)
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

	tags, err := c.questionsService.GetTagsForQuestion(context, intId)
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
		return
	}

	questionId, _ := strconv.ParseInt(request.ID, 10, 0)

	result, err := c.questionsService.GetAllSubmissionsForQuestion(
		context,
		int(questionId),
	)
	if err != nil {
		context.JSON(500, gin.H{
			"error": "An internal server error occurred.",
		})
		return
	}

	resp := models.Pagaination[models.QuestionSubmission]{
		Data: result,
	}

	if len(result) == 0 {
		resp.Data = make([]models.QuestionSubmission, 0)
	}

	context.JSON(200, resp)
}

func (c QuestionsController) GetAllQuestionTags(context *gin.Context) {
	tags, err := c.questionsService.GetAllQuestionTags(context)
	if err != nil {
		context.JSON(500, gin.H{
			"error": "An internal server error has occurred.",
		})
		return
	}

	context.JSON(200, gin.H{
		"tags": tags,
	})
}

func (c QuestionsController) SaveQuestionSubmission(context *gin.Context) {
	var questionSubmissionRequest saveQuestionSubmissionRequest
	if err := context.ShouldBindBodyWithJSON(&questionSubmissionRequest); err != nil {
		context.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := c.questionsService.SaveQuestionSubmission(
		context,
		questionSubmissionRequest.QuestionID,
		uuid.New(),
		time.Now(),
		time.Duration(questionSubmissionRequest.TimeTaken*int(time.Second)),
		models.ConfidenceLevel(questionSubmissionRequest.ConfidenceLevel),
	); err != nil {
		context.JSON(500, gin.H{
			"error": "An internal error has occurred.",
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Successfully saved question submission",
	})
}
