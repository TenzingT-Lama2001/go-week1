package main

import (
	"log"
	"week1/app/controller"
	pb "week1/app/gen"
	"week1/app/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	// Start HTTP server
	startHTTPServer()
}

func startHTTPServer() {
	// Create a gRPC client
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	blogServiceClient := pb.NewBlogServiceClient(conn)

	blogController := controller.NewBlogController(blogServiceClient)

	// Create a new gin router
	router := gin.Default()

	// Setup routes for the blog API
	routes.SetupBlogRoutes(router, blogController)

	// Start the HTTP server
	err = router.Run(":8081")
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
