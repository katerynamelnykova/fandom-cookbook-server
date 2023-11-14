package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortFandomInfo struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Fandom     string             `json:"fandom" bson:"fandom"`
	BgImage    string             `json:"bgImage" bson:"bgImage"`
	FrontImage string             `json:"frontImage" bson:"frontImage"`
	Logo       string             `json:"logo" bson:"logo"`
}
