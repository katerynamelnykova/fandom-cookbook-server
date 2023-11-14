package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Fandom      string             `json:"fandom" bson:"fandom"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Recipes     []RecipeNameId     `json:"recipes" bson:"recipes"`
	BgColor     string             `json:"bgColor" bson:"bgColor"`
	FrontColor  string             `json:"frontColor" bson:"frontColor"`
	TextColor   string             `json:"textColor" bson:"textColor"`
	Logo        string             `json:"logo" bson:"logo"`
	FrontImage  string             `json:"frontImage" bson:"frontImage"`
	BgImage     string             `json:"bgImage" bson:"bgImage"`
	BookImage   string             `json:"bookImage" bson:"bookImage"`
}

type FullBook struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Fandom      string             `json:"fandom" bson:"fandom"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Recipes     []Recipe           `json:"recipes" bson:"recipes"`
	BgColor     string             `json:"bgColor" bson:"bgColor"`
	FrontColor  string             `json:"frontColor" bson:"frontColor"`
	TextColor   string             `json:"textColor" bson:"textColor"`
	Logo        string             `json:"logo" bson:"logo"`
	FrontImage  string             `json:"frontImage" bson:"frontImage"`
	BgImage     string             `json:"bgImage" bson:"bgImage"`
	BookImage   string             `json:"bookImage" bson:"bookImage"`
}
