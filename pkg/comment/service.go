package comment

import (
	"github.com/jordanpagewhite/go-fiber-first/api/presenter"
	"github.com/jordanpagewhite/go-fiber-first/pkg/entities"
)

type Service interface {
	InsertComment(comment *entities.Comment) (*entities.Comment, error)
	FetchComments() (*[]presenter.Comment, error)
	UpdateComment(comment *entities.Comment) (*entities.Comment, error)
	RemoveComment(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertComment(comment *entities.Comment) (*entities.Comment, error) {
	return s.repository.CreateComment(comment)
}

func (s *service) FetchComments() (*[]presenter.Comment, error) {
	return s.repository.ReadComment()
}

func (s *service) UpdateComment(comment *entities.Comment) (*entities.Comment, error) {
	return s.repository.UpdateComment(comment)
}

func (s *service) RemoveComment(ID string) error {
	return s.repository.DeleteComment(ID)
}
