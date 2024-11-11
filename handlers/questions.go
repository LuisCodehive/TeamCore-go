package handlers

import (
	"fmt"
	"net/http"
	"teamcore_service/services"
	"time"

	"github.com/gin-gonic/gin"
)

func GetFormattedQuestions(c *gin.Context) {
	questions, err := services.FetchQuestions()

	fmt.Println("Error fetching questions:", err)
	fmt.Println("questions:", questions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GetFormattedQuestions Failed to fetch questions"})
		return
	}

	formattedQuestions := make([]gin.H, len(questions))
	for i, q := range questions {
		formattedQuestions[i] = gin.H{"pregunta_id": q.QuestionID, "pregunta": q.Question, "respuestas": q.Answers}
	}

	c.JSON(http.StatusOK, gin.H{
		"titulo":      "Preguntas del día",
		"dia":         time.Now().Format("02-01-2006"),
		"info":        formattedQuestions,
		"api_version": 1,
	})
}
