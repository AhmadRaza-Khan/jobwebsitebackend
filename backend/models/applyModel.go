package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Apply struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Frontend    string             `json:"frontend" bson:"frontend"`
	Backend     string             `json:"backend" bson:"backend"`
	DevOps      string             `json:"devOps" bson:"devOps"`
	Databases   string             `json:"databases" bson:"databases"`
	Cloud       string             `json:"cloud" bson:"cloud"`
	Engineering string             `json:"engineering" bson:"engineering"`
	Experience  string             `json:"experience" bson:"experience"`
}
type Status struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Frontend    string `json:"frontend"`
	Backend     string `json:"backend"`
	DevOps      string `json:"devOps"`
	Databases   string `json:"databases"`
	Cloud       string `json:"cloud"`
	Engineering string `json:"engineering"`
	Experience  string `json:"experience"`
}
type Check struct {
	Email string `json:"email" bson:"email"`
}
