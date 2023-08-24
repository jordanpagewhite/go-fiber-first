package main

import (
	"context"
	"fmt"
	"log"
	"os"
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
	commentCollection := db.Collection(os.Getenv("DB_NAME"))
	commentRepo := comment.NewRepo(commentCollection)
	commentService := comment.NewService(commentRepo)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")
	routes.CommentRouter(v1, commentService)

	defer cancel()
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db_string := fmt.Sprintf("mongodb://%s:%s/%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		db_string).SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database(os.Getenv("DB_NAME"))
	return db, cancel, nil
}
