package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortFandomInfo struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Fandom     string             `json:"fandom" bson:"fandom"`
	Logo       string             `json:"logo" bson:"logo"`
	FrontImage string             `json:"frontImage" bson:"frontImage"`
	BgImage    string             `json:"bgImage" bson:"bgImage"`
}
