package presenter

import (
	"time"

	"github.com/jordanpagewhite/go-fiber-first/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`                // this comment's ID
	ParentID  primitive.ObjectID `json:"parent_id" bson:"parent_id"`   // parent comment ID
	UserID    string             `json:"user_id" bson:"user_id"`       // SSO user_id
	Subject   string             `json:"subject" bson:"subject"`       // The subject of the comment
	Body      string             `json:"body" bson:"body"`             // The body of the comment
	CreatedAt time.Time          `json:"created_at" bson:"created_at"` // The UNIX timestamp when the comment was created.
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"` // The UNIX timestamp when the comment was last changed.
	Private   bool               `json:"private" bson:"private"`       // Whether or not this is a private comment.
}

func CommentSuccessResponse(data *entities.Comment) *fiber.Map {
	comment := Comment{
		ID:        data.ID,
		ParentID:  data.ParentID,
		UserID:    data.UserID,
		Subject:   data.Subject,
		Body:      data.Body,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Private:   data.Private,
	}
	return &fiber.Map{
		"status": true,
		"data":   comment,
		"error":  nil,
	}
}

func CommentsSuccessResponse(data *[]Comment) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func CommentErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
