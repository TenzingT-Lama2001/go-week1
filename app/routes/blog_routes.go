package routes

import (
	"week1/app/controller"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the blog controller.
func SetupBlogRoutes(router *gin.Engine, controller *controller.BlogController) {
	router.GET("/posts", controller.GetPosts)
	router.GET("/posts/:id", controller.GetPost)
	router.POST("/posts", controller.CreatePost)
	router.DELETE("/posts/:id", controller.DeletePost)
	router.PATCH("/posts/:id", controller.UpdatePost)
}
