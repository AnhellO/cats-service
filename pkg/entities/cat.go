package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cat struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Breed string             `bson:"breed" json:"breed"`
	Age   int                `bson:"age" json:"age"`
	Image string             `bson:"image" json:"image"`
}
