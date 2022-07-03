package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title string             `json:"title,omitempty" form:"title" validate:"required" bson:"title"`
	Image string             `json:"image,omitempty" bson:"image"`

	// Description string             `json:"description" bson:"description"`
	// Option_A    string             `json:"option_a" bson:"option_a"`
	// Option_B    string             `json:"option_b" bson:"option_b"`
	// Option_C    string             `json:"option_c" bson:"option_c"`
	// Option_D    string             `json:"option_d" bson:"option_d"`
	// Option_E    string             `json:"option_e" bson:"option_e"`
	// Answer      string             `json:"answer" bson:"answer"`
	// Category    string             `json:"category" bson:"category"`
	// UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Difficulty string `json:"difficulty" bson:"difficulty"`
	// CreatedAt   string             `json:"created_at" bson:"created_at"`
	// UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
}
