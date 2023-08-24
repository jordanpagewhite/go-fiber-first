package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jordanpagewhite/go-fiber-first/api/routes"
	"github.com/jordanpagewhite/go-fiber-first/pkg/comment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	commentCollection := db.Collection("comments")
	commentRepo := comment.NewRepo(commentCollection)
	commentService := comment.NewService(commentRepo)

	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	routes.CommentRouter(v1, commentService)

	defer cancel()
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://username:password@localhost:27017/fiber").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("comments")
	return db, cancel, nil
}
