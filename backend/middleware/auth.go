package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ahmadraza-khan/jobwebsite/config"
	"github.com/ahmadraza-khan/jobwebsite/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Authentication(c *gin.Context) {
	// Get the JWT token from the cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		// Log error if Authorization cookie is missing
		log.Printf("ERROR: Authorization cookie missing: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie missing"})
		return
	}
	log.Printf("INFO: Received token for authentication")

	// Validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("ERROR: Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		log.Printf("INFO: Token signed with valid method: %v", token.Method)
		return []byte(os.Getenv("JwtSecret")), nil
	})

	// Handle invalid token
	if err != nil || !token.Valid {
		log.Printf("ERROR: Invalid token: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["exp"] == nil || claims["sub"] == nil {
		log.Printf("ERROR: Invalid token claims")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Check token expiration
	exp := int64(claims["exp"].(float64)) // Convert expiration to int64
	if time.Now().Unix() > exp {
		log.Printf("ERROR: Token expired")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}
	log.Printf("INFO: Token is valid and not expired")

	// Convert claims["sub"] to an ObjectId
	userID, err := primitive.ObjectIDFromHex(claims["sub"].(string))
	if err != nil {
		log.Printf("ERROR: Invalid user ID in token: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	// Fetch user from the database
	var user models.User
	userDB := config.GetCollection("userData")
	err = userDB.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Printf("ERROR: User not found in database: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	log.Printf("INFO: User %s authenticated successfully", user.Id.Hex())

	// Set user data in context and proceed
	c.Set("user", user)
	c.Next()
}
