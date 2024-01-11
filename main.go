package main

import (
	"log"
	"net"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"week1/app/blog/api"
	"week1/app/blog/pb"
	"week1/app/blog/service"
)

// App struct to hold dependencies like the database
type App struct {
	DB *sql.DB
}

func main() {
	db, err := sql.Open("mysql", "root:Password!?@#$123@tcp(localhost:3306)/blogs")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

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
