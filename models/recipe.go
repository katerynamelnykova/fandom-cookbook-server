package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecipeNameId struct {
	ID   primitive.ObjectID `json:"_idrec" bson:"_idrec"`
	Name string             `json:"name" bson:"name"`
}

type Recipe struct {
	ID        primitive.ObjectID `json:"_idrec" bson:"_idrec"`
	Name      string             `json:"name" bson:"name"`
	Recipe    string             `json:"recipe" bson:"recipe"`
	VideoLink string             `json:"videoLink" bson:"videoLink"`
}
