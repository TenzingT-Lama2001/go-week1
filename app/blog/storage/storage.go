// storage.go
package storage

import (
	"github.com/google/uuid"
)

// Post represents data about a blog post.
type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Posts slice to seed blog post data.
var Posts = []Post{
	{ID: "1", Title: "Introduction to Go", Content: "This is a blog post about Go programming language."},
	{ID: "2", Title: "Building RESTful APIs with Gin", Content: "Learn how to build RESTful APIs using the Gin framework in Go."},
	{ID: "3", Title: "Concurrency in Go", Content: "Explore the concurrency features of the Go programming language."},
}

// GenerateUniqueID generates a unique ID for blog posts.
func GenerateUniqueID() string {
	return uuid.New().String()
}
