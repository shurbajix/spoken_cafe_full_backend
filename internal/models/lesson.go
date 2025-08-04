package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lesson struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	TeacherID   primitive.ObjectID `bson:"teacherId" json:"teacherId"`
	StudentID   primitive.ObjectID `bson:"studentId" json:"studentId"`
	Language    string             `bson:"language" json:"language"`
	Level       string             `bson:"level" json:"level"`
	Duration    int                `bson:"duration" json:"duration"`
	Price       float64            `bson:"price" json:"price"`
	Status      string             `bson:"status" json:"status"`
	Schedule    time.Time          `bson:"schedule" json:"schedule"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
