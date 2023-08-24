package comment

import (
	"context"
	"time"

	"github.com/jordanpagewhite/go-fiber-first/api/presenter"
	"github.com/jordanpagewhite/go-fiber-first/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateComment(comment *entities.Comment) (*entities.Comment, error)
	ReadComment() (*[]presenter.Comment, error)
	UpdateComment(comment *entities.Comment) (*entities.Comment, error)
	DeleteComment(ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

// CreateComment is a mongo repository that helps to create comments
func (r *repository) CreateComment(comment *entities.Comment) (*entities.Comment, error) {
	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// ReadComment is a mongo repository that helps to fetch comments
func (r *repository) ReadComment() (*[]presenter.Comment, error) {
	var comments []presenter.Comment
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var comment presenter.Comment
		_ = cursor.Decode(&comment)
		comments = append(comments, comment)
	}
	return &comments, nil
}

// UpdateComment is a mongo repository that helps to update comments
func (r *repository) UpdateComment(comment *entities.Comment) (*entities.Comment, error) {
	comment.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": comment.ID}, bson.M{"$set": comment})
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// DeleteComment is a mongo repository that helps to delete comments
func (r *repository) DeleteComment(ID string) error {
	commentID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": commentID})
	if err != nil {
		return err
	}
	return nil
}
