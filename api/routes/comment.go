package routes

import (
	"github.com/jordanpagewhite/go-fiber-first/api/handlers"
	"github.com/jordanpagewhite/go-fiber-first/pkg/comment"

	"github.com/gofiber/fiber/v2"
)

func CommentRouter(app fiber.Router, service comment.Service) {
	app.Get("/comment", handlers.GetComment(service))
	app.Post("/comment", handlers.AddComment(service))
	app.Put("/comment", handlers.UpdateComment(service))
	app.Delete("/comment", handlers.RemoveComment(service))
}
