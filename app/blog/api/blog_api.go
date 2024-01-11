package api

import (
	"context"
	"net/http"
	"week1/app/blog/pb"
	"fmt"
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

// SetupRoutes sets up the routes for the blog controller.
func (controller *BlogController) SetupRoutes(router *gin.Engine) {
	router.GET("/posts", controller.getPosts)
	router.GET("/posts/:id", controller.getPostByID)
	router.POST("/posts", controller.createPost)
	router.DELETE("/posts/:id", controller.deletePostByID)
	router.PATCH("/posts/:id", controller.updatePostByID)
}


func (controller *BlogController) getPosts(c *gin.Context) {
	// Use the gRPC client to call the GetPosts method on the gRPC server
	response, err := controller.BlogService.GetPosts(context.Background(), &pb.Empty{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get posts from gRPC service"})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (controller *BlogController) getPostByID(c *gin.Context) {
	id := c.Param("id")

	// Use the gRPC client to call the GetPost method on the gRPC server
	response, err := controller.BlogService.GetPost(context.Background(), &pb.GetPostRequest{Id: id})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (controller *BlogController) createPost(c *gin.Context) {
	var newPost pb.CreatePostRequest
	if err := c.BindJSON(&newPost); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid JSON"})
		return
	}
	fmt.Println("newPost", &newPost)

	// Use the gRPC client to call the CreatePost method on the gRPC server
	response, err := controller.BlogService.CreatePost(context.Background(), &newPost)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to create post"})
		return
	}

	c.IndentedJSON(http.StatusCreated, response)
}

func (controller *BlogController) deletePostByID(c *gin.Context) {
	id := c.Param("id")

	// Use the gRPC client to call the DeletePost method on the gRPC server
	_, err := controller.BlogService.DeletePost(context.Background(), &pb.GetPostRequest{Id: id})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func (controller *BlogController) updatePostByID(c *gin.Context){
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