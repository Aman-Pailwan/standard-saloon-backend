package api

import (
	"log"
	"net/http"
	"os"

	"github.com/Aman-Pailwan/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env config: ", err)
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate schema: ", err)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		models.ResponseJson(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	DB.Create(&user)
	models.ResponseJson(c, http.StatusCreated, "User Created Successfully", user)
}
