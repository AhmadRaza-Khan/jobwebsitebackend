package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name required"`
	Email     string             `json:"email" bson:"email"`
	Phone     string             `json:"phone" bson:"phone"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64              `json:"updatedAt" bson:"updatedAt"`
	Token     string             `json:"token" bson:"token"`
	UserType  string             `json:"userType" bson:"userType"`
}
type ClientUser struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
	UserType string `json:"userType" bson:"userType"`
}
type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
type Verification struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Otp       int64              `json:"otp" bson:"otp"`
	Status    bool               `json:"status" bson:"status"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}
