package models

import "github.com/gin-gonic/gin"

type User struct {
	FullName       string        `json:"full_name"`
	Email          string        `json:"email"`
	PhoneNumber    int           `json:"phone_number"`
	PrimaryService []interface{} `json:"primary_service"`
}

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseJson(c *gin.Context, status int, message string, data any) {
	response := JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.JSON(status, response)
}
