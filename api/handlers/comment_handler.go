package handlers

import (
	"github.com/jordanpagewhite/go-fiber-first/api/presenter"
	"github.com/jordanpagewhite/go-fiber-first/pkg/comment"
	"github.com/jordanpagewhite/go-fiber-first/pkg/entities"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func AddComment(service comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Comment
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.CommentErrorResponse(err))
		}
		if requestBody.Subject == "" || requestBody.Body == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CommentErrorResponse(errors.New(
				"Please specify a subject and body")))
		}
		result, err := service.InsertComment(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CommentErrorResponse(err))
		}
		return c.JSON(presenter.CommentSuccessResponse(result))
	}
}

func UpdateComment(service comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Comment
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.CommentErrorResponse(err))
		}
		result, err := service.UpdateComment(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CommentErrorResponse(err))
		}
		return c.JSON(presenter.CommentSuccessResponse(result))
	}
}

func RemoveComment(service comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.CommentErrorResponse(err))
		}
		commentID := requestBody.ID
		err = service.RemoveComment(commentID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CommentErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetComment(service comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchComments()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CommentErrorResponse(err))
		}
		return c.JSON(presenter.CommentsSuccessResponse(fetched))
	}
}
