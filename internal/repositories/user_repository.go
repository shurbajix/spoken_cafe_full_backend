package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/spokencafe/backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRespository(db *mongo.Database) *UserRepository {
	repo := &UserRepository{
		collection: db.Collection("users"),
	}
	// Create unique index on username/email for fast lookups & uniqueness

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
	return repo
}

func (r *UserRepository) createUser(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return &user, err
}
