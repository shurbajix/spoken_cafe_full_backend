package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	LessonID  primitive.ObjectID `bson:"lessonId" json:"lessonId"`
	StudentID primitive.ObjectID `bson:"`
}
