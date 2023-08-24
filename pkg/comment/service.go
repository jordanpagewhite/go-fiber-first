package comment

import (
	"github.com/jordanpagewhite/go-fiber-first/api/presenter"
	"github.com/jordanpagewhite/go-fiber-first/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertComment(comment *entities.Comment) (*entities.Comment, error)
	FetchComments() (*[]presenter.Comment, error)
	UpdateComment(comment *entities.Comment) (*entities.Comment, error)
	RemoveComment(ID string) error
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// InsertComment is a service layer that helps insert comment in CommentShop
func (s *service) InsertComment(comment *entities.Comment) (*entities.Comment, error) {
	return s.repository.CreateComment(comment)
}

// FetchComments is a service layer that helps fetch all comments in CommentShop
func (s *service) FetchComments() (*[]presenter.Comment, error) {
	return s.repository.ReadComment()
}

// UpdateComment is a service layer that helps update comments in CommentShop
func (s *service) UpdateComment(comment *entities.Comment) (*entities.Comment, error) {
	return s.repository.UpdateComment(comment)
}

// RemoveComment is a service layer that helps remove comments from CommentShop
func (s *service) RemoveComment(ID string) error {
	return s.repository.DeleteComment(ID)
}
