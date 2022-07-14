package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Course      string             `json:"course" bson:"course,omitempty" validate:"required,min=3"`
	Image       string             `json:"image" bson:"image,omitempty"`
	Description string             `json:"description" bson:"description,omitempty" validate:"required"`
	Subject     string             `json:"subject" bson:"subject,omitempty" validate:"required"`
	Difficulty  string             `json:"difficulty" bson:"difficulty,omitempty" validate:"required"`

	Option_A string `json:"option_a" bson:"option_a,omitempty" validate:"required"`
	Option_B string `json:"option_b" bson:"option_b,omitempty" validate:"required"`
	Option_C string `json:"option_c" bson:"option_c,omitempty" validate:"required"`
	Option_D string `json:"option_d" bson:"option_d,omitempty" validate:"required"`
	Option_E string `json:"option_e" bson:"option_e,omitempty" validate:"required"`
	Answer   string `json:"answer" bson:"answer,omitempty"`

	IsActive bool `json:"isActive" bson:"isActive"`
}
