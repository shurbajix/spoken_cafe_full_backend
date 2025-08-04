package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username          string             `bson:"username" json:"username"`
	Password          string             `bson:"password" json:"_"`
	Email             string             `bson:"email" json:"email"`
	Role              string             `bson:"role" json:"role"`
	CreatedAt         time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time          `bson:"updatedAt" json:"updatedAt"`
	DeletedAt         time.Time          `bson:"deletedAt" json:"deletedAt"`
	IsActive          bool               `bson:"isActive" json:"isActive"`
	IsDeleted         bool               `bson:"isDeleted" json:"isDeleted"`
	IsVerified        bool               `bson:"isVerified" json:"isVerified"`
	IsAdmin           bool               `bson:"isAdmin" json:"isAdmin"`
	IsTeacher         bool               `bson:"isTeacher" json:"isTeacher"`
	IsStudent         bool               `bson:"isStudent" json:"isStudent"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	ProfilePicture    string             `bson:"profilePicture" json:"profilePicture"`
	Bio               string             `bson:"bio" json:"bio"`
	Language          string             `bson:"language" json:"language"`
	LearningLanguages []string           `bson:"learningLanguages" json:"learningLanguages"`
	Phone             string             `bson:"phone" json:"phone"`
	DateOfBirth       time.Time          `bson:"dateOfBirth" json:"dateOfBirth"`
	Gender            string             `bson:"gender" json:"gender"`
}
