package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type Answer struct {
	AnswerID string `json:"answer_id"`
	Answer   string `json:"answer"`
}

type Question struct {
	QuestionID string   `json:"question_id"`
	Question   string   `json:"question"`
	Answers    []Answer `json:"answers"`
}

type QuestionsResponse struct {
	Date string     `json:"date"`
	Data []Question `json:"data"`
}

func FetchQuestions() ([]Question, error) {
	client := resty.New()
	authToken := os.Getenv("AUTH_TOKEN")
	print(authToken)
	resp, err := client.R().
		SetHeader("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NzM0NzU4MTEsImV4cCI6MTcwNTAxMTgxMSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.9wqriO_2Q8Xfwc9VcgMpr-2c4WVdLRJ5G6NcNaXdpuY").
		Get("https://us-central1-teamcore-retail.cloudfunctions.net/test_mobile/api/questions")

	if err != nil || resp.StatusCode() != 200 {
		return nil, errors.New("FetchQuestions failed to fetch questions")
	}

	var questionsResp QuestionsResponse

	// Imprime el contenido de questionsResp para verificar los datos
	fmt.Printf("Questions Response: %+v\n", questionsResp)

	// Verifica que las respuestas estén siendo deserializadas correctamente
	for _, question := range questionsResp.Data {
		fmt.Printf("Question: %+v\n", question)
		for _, answer := range question.Answers {
			fmt.Printf("Answer: %+v\n", answer)
		}
	}

	err = json.Unmarshal(resp.Body(), &questionsResp)
	if err != nil {
		return nil, err
	}

	return questionsResp.Data, nil
}
