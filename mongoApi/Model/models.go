package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotstar struct {
	ID  primitive.ObjectID `json:"_id,omitempty bson:"_id,omitempty"`
	Movie string `json:"movie,omitempty"`
	Series string `json:"series,omitempty"`
	Watched bool `json:"watched,omitempty"`
}