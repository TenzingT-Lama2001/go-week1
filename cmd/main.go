package main

import (
	"database/sql"
	"log"
	"net"
	"week1/app/controller"
	"week1/app/database"
	pb "week1/app/gen"
	"week1/app/routes"
	"week1/app/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

// App struct to hold dependencies like the database
type App struct {
	DB *sql.DB
}

func main() {

	// Connect to the database
	db := database.NewDB()

	// Close the database connection when the main function finishes
	defer database.CloseDB(db)

	// Create an instance of the App struct
	app := &App{DB: db}

	// Start gRPC server
	go startGRPCServer(app)

	// Start HTTP server
	startHTTPServer()
}

func startGRPCServer(app *App) {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverRegistry := grpc.NewServer()
	blogGRPCService := service.NewBlogGRPCService(app.DB)
	pb.RegisterBlogServiceServer(serverRegistry, blogGRPCService)

	if err := serverRegistry.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
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
