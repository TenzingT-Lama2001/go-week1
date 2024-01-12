package controller

import (
	"context"
	"fmt"
	"net/http"
	pb "week1/app/gen"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogService pb.BlogServiceClient
}

func NewBlogController(blogService pb.BlogServiceClient) *BlogController {
	return &BlogController{
		BlogService: blogService,
	}
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (controller *BlogController) GetPosts(c *gin.Context) {
	// Use the gRPC client to call the GetPosts method on the gRPC server
	response, err := controller.BlogService.GetPosts(context.Background(), &pb.Empty{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get posts from gRPC service"})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (controller *BlogController) GetPost(c *gin.Context) {
	id := c.Param("id")

	// Use the gRPC client to call the GetPost method on the gRPC server
	response, err := controller.BlogService.GetPost(context.Background(), &pb.GetPostRequest{Id: id})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (controller *BlogController) CreatePost(c *gin.Context) {
	var newPost pb.CreatePostRequest
	if err := c.BindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid JSON"})
		return
	}

	if newPost.Title == "" || newPost.Content == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please check your inputs"})
		return
	}
	fmt.Println("newPost", &newPost)

	//Check if the post exists in database
	exists, err := controller.BlogService.PostExists(context.Background(), &pb.PostExistsRequest{Title: newPost.Title})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to check if post exists"})
		return
	}

	if exists.Exists {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "post already exists"})
		return
	}

	// Use the gRPC client to call the CreatePost method on the gRPC server
	response, err := controller.BlogService.CreatePost(context.Background(), &newPost)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to create post"})
		return
	}

	c.IndentedJSON(http.StatusCreated, response)
}

func (controller *BlogController) DeletePost(c *gin.Context) {
	id := c.Param("id")

	// Use the gRPC client to call the DeletePost method on the gRPC server
	_, err := controller.BlogService.DeletePost(context.Background(), &pb.GetPostRequest{Id: id})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func (controller *BlogController) UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var updatePost pb.UpdatePostRequest
	if err := c.BindJSON(&updatePost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid JSON"})
		return
	}

	// Use the gRPC client to call the UpdatePost method on the gRPC server
	response, err := controller.BlogService.UpdatePost(context.Background(), &pb.UpdatePostRequest{Id: id, Title: updatePost.Title, Content: updatePost.Content})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to update post"})
		return
	}

	c.IndentedJSON(http.StatusOK, response)

}
