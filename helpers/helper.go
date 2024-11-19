package helpers

import (
	"errors"
	"log"
	"net/http"

	"github.com/ahmadraza-khan/jobwebsite/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserValidation(data *models.ClientUser) error {
	if data.Name == "" {
		return errors.New("name can't be empty")
	}
	if data.Email == "" {
		return errors.New("name can't be empty")
	}
	if data.Phone == "" {
		return errors.New("name can't be empty")
	}
	if data.Password == "" {
		return errors.New("name can't be empty")
	}
	return nil
}
func ApplyValidation(data *models.Apply) error {
	if data.Name == "" {
		return errors.New("name can't be empty")
	}
	if data.Email == "" {
		return errors.New("email can't be empty")
	}
	if data.Backend == "" {
		return errors.New("backend can't be empty")
	}
	if data.Frontend == "" {
		return errors.New("frontend can't be empty")
	}
	if data.Cloud == "" {
		return errors.New("cloud computing can't be empty")
	}
	if data.Databases == "" {
		return errors.New("database can't be empty")
	}
	if data.Engineering == "" {
		return errors.New("engineering can't be empty")
	}
	if data.DevOps == "" {
		return errors.New("devOps can't be empty")
	}
	if data.Experience == "" {
		return errors.New("experience can't be empty")
	}
	return nil
}

func CheckError(c *gin.Context, err error) {
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(bytes)
}
