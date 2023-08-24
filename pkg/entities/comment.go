package entities

import (
	"time"

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

type Comments struct {
	Pages          int       `json:"pages" bson:"pages"`                   // The total number of pages resulting from this comment thread and the provided page_size
	Previous_page  int       `json:"previous_page" bson:"previous_page"`   // The 0-indexed integer representing the previous page OR NULL if current_page is 0 and there is no previous page.
	Current_page   int       `json:"current_page" bson:"current_page"`     // The 0-indexed integer representing the current page
	Next_page      int       `json:"next_page" bson:"next_page"`           // The 0-indexed integer representing the next page OR NULL if current_page is the final page and there is no next page.
	Page_size      int       `json:"page_size" bson:"page_size"`           // The maximum number of comments to include per page.
	Source_abs_url string    `json:"source_abs_url" bson:"source_abs_url"` // The absolute URL of the source page of this comment thread
	Comment_count  int       `json:"comment_count" bson:"comment_count"`   // How many total comments are there in this thread? This also is used to determine the total number of comments for pagination.
	Comment_latest int       `json:"comment_latest" bson:"comment_latest"` // If this thread has one or more comments, this returns the UNIX timestamp of when the most recently updated comment was changed
	Comments       []Comment `json:"comments" bson:"comments"`             // The comments in this comment thread.
}

type DeleteRequest struct {
	ID string `json:"id" bson:"_id"`
}
