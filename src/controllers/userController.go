package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/ahmadraza-khan/jobwebsite/src/config"
	"github.com/ahmadraza-khan/jobwebsite/src/helpers"
	"github.com/ahmadraza-khan/jobwebsite/src/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.ClientUser
	var dbUser models.User

	if err := c.BindJSON(&user); err != nil {
		helpers.CheckError(c, err)
		return
	}

	if err := helpers.UserValidation(&user); err != nil {
		helpers.CheckError(c, err)
		return
	}
	filter := bson.M{"email": user.Email}
	userDB := config.GetCollection("userData")
	result := userDB.FindOne(context.TODO(), filter)
	if result.Err() == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "the email is already registered"})
		return
	} else if result.Err() != mongo.ErrNoDocuments {
		helpers.CheckError(c, result.Err())
		return
	}

	hashedPass := helpers.HashPassword(user.Password)

	dbUser.Name = user.Name
	dbUser.Email = user.Email
	dbUser.Phone = user.Phone
	dbUser.Password = hashedPass
	dbUser.CreatedAt = time.Now().Unix()
	dbUser.UpdatedAt = time.Now().Unix()
	dbUser.UserType = "GENERAL"

	insertResult, err := userDB.InsertOne(context.TODO(), dbUser)
	if err != nil {
		helpers.CheckError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "success": true, "message": "user created successfully", "id": insertResult.InsertedID})
}

func Login(c *gin.Context) {
	var user models.Login
	var dbUser models.User

	err := c.BindJSON(&user)
	if err != nil {
		helpers.CheckError(c, err)
		return
	}

	// Fetch user from database
	filter := bson.M{"email": user.Email}
	userDB := config.GetCollection("userData")
	result := userDB.FindOne(context.TODO(), filter)

	if result.Err() != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Email is not registered"})
		return
	}

	err = result.Decode(&dbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to retrieve user data"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Incorrect password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JwtSecret")))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Failed to create token"})
		return
	}

	c.SetCookie("Authorization", tokenString, 3600*24, "/", "", true, true)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Login successful",
		"token":    tokenString,
		"error":    false,
		"userType": dbUser.UserType,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"success": true,
		"message": "Logout successful",
	})
}
