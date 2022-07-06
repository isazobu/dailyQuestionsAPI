package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateQuestion struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" form:"title" validate:"required" bson:"title"`
	Image       string             `json:"image" bson:"image"`
	Description string             `json:"description" bson:"description"`
	Category    string             `json:"category" bson:"category"`
	Difficulty  string             `json:"difficulty" bson:"difficulty"`

	Option_A string `json:"option_a" bson:"option_a"`
	Option_B string `json:"option_b" bson:"option_b"`
	Option_C string `json:"option_c" bson:"option_c"`
	Option_D string `json:"option_d" bson:"option_d"`
	Option_E string `json:"option_e" bson:"option_e"`
	Answer   string `json:"answer" bson:"answer"`

	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
