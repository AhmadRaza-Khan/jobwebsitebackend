package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/ahmadraza-khan/jobwebsite/src/config"
	"github.com/ahmadraza-khan/jobwebsite/src/helpers"
	"github.com/ahmadraza-khan/jobwebsite/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Apply(c *gin.Context) {
	var apply models.Apply

	// Get the "applyData" collection from MongoDB
	dbApply := config.GetCollection("applyData")

	// Bind the incoming JSON data to the 'apply' struct
	err := c.BindJSON(&apply)
	if err != nil {
		helpers.CheckError(c, err)
		return
	}

	// Validate the application data
	err = helpers.ApplyValidation(&apply)
	if err != nil {
		helpers.CheckError(c, err)
		return
	}

	// Insert the application data into MongoDB
	result, err := dbApply.InsertOne(context.TODO(), apply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error inserting data into database",
			"error":   err.Error(),
		})
		return
	}

	// Respond with the result of the insertion
	c.JSON(http.StatusOK, gin.H{
		"message": "Application submitted successfully!",
		"data":    result,
		"success": true,
	})
}

func GetApplicationStatus(c *gin.Context) {
	var user models.Check

	// Bind the incoming JSON data
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Failed to parse JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
		})
		return
	}

	// Validate email
	if user.Email == "" {
		log.Println("Email field is empty")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is required"})
		return
	}

	// Get the "applyData" collection
	dbApply := config.GetCollection("applyData")
	var application models.Status

	// Query the database for the application by email
	err := dbApply.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&application)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("Application not found for email:", user.Email)
			c.JSON(http.StatusNotFound, gin.H{"message": "Application not found"})
		} else {
			log.Println("Database error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}
		return
	}

	// Log the successful retrieval of the application
	log.Println("Application found for email:", user.Email)

	// Return the application data
	c.JSON(http.StatusOK, gin.H{
		"data": application,
	})
}
