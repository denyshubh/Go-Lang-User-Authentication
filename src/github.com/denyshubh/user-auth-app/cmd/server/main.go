package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/denyshubh/user-auth-app/internal/graph"
	"github.com/denyshubh/user-auth-app/internal/graph/generated"
	"github.com/denyshubh/user-auth-app/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Run migrations
	err = db.AutoMigrate(&model.User{}, &model.Authentication{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Set up Gin router
	router := gin.Default()

	// GraphQL endpoint
	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(db)}))
	router.POST("/graphql", func(c *gin.Context) {
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	// Start the server
	port := 8080
	fmt.Printf("Server running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func rest() {
	// Initialize the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&User{}, &Authentication{})

	// Set up Gin router
	router := gin.Default()

	// REST API endpoints
	router.POST("/users", func(c *gin.Context) {
		// Create user
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		// Update user
	})

	router.GET("/users/:id", func(c *gin.Context) {
		// Get user
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		// Delete user
	})

	router.PUT("/users/:id/password", func(c *gin.Context) {
		// Update password
	})

	// Start the server
	router.Run()
}