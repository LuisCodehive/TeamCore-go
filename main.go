package main

import (
    "teamcore_service/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/formatted-questions", handlers.GetFormattedQuestions)
    router.Run(":8080")
}
