package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" form:"title" validate:"required" bson:"title,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Category    string             `json:"category" bson:"category,omitempty"`
	Difficulty  string             `json:"difficulty" bson:"difficulty,omitempty"`

	Option_A string `json:"option_a" bson:"option_a,omitempty"`
	Option_B string `json:"option_b" bson:"option_b,omitempty"`
	Option_C string `json:"option_c" bson:"option_c,omitempty"`
	Option_D string `json:"option_d" bson:"option_d,omitempty"`
	Option_E string `json:"option_e" bson:"option_e,omitempty"`
	Answer   string `json:"answer" bson:"answer,omitempty"`
}
