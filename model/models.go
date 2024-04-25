package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Theatre struct {
	ID      primitive.ObjectID `json:"_id,omitemppty" bson:"_id,omitempty"`
	Watched bool               `json:"watched,omitempty"`
	Movie   string             `json:"movie,omitempty"`
}
