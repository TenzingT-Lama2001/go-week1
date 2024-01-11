// main.go

package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"week1/app/blog/api"
	"week1/app/blog/pb"
	"week1/app/blog/service"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverRegistry := grpc.NewServer()
	blogGRPCService := service.NewBlogGRPCService()
	pb.RegisterBlogServiceServer(serverRegistry, blogGRPCService)

	go func() {
		if err := serverRegistry.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()

	// Create a gRPC client
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	blogServiceClient := pb.NewBlogServiceClient(conn)

	// Create the API and inject the gRPC client
	blogAPI := api.NewBlogController(blogServiceClient)

	// Create a new gin router
	router := gin.Default()

	// Setup routes for the blog API
	blogAPI.SetupRoutes(router)

	// Start the HTTP server
	err = router.Run(":8081")
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
