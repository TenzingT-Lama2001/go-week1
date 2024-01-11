// service/blog_grpc.go
package service

import (
	"context"
	"fmt"
	"week1/app/blog/pb"
	"github.com/google/uuid"
    "database/sql"
)

type BlogGRPCService struct{
	pb.UnimplementedBlogServiceServer
    DB *sql.DB
}

func NewBlogGRPCService(db  *sql.DB) *BlogGRPCService {
	return &BlogGRPCService{DB: db}
}

func (s *BlogGRPCService) GetPosts(ctx context.Context, req *pb.Empty) (*pb.PostList, error) {
    blogs, err := s.DB.Query("SELECT * FROM blog")
    if err != nil {
        fmt.Println(err)
    }
    defer blogs.Close()

    fmt.Println("blogs", blogs)
    
    var postList []*pb.Post
    for blogs.Next() {
        var post pb.Post
        err := blogs.Scan(&post.Id, &post.Title, &post.Content)
        if err != nil {
            fmt.Println(err)
        }
        postList = append(postList, &post)
    }
    return &pb.PostList{Posts: postList}, nil

}


func (s *BlogGRPCService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
    id := req.Id

    blog := s.DB.QueryRow("SELECT * FROM blog WHERE id = ?", id) 
    var post pb.Post
    err := blog.Scan(&post.Id, &post.Title, &post.Content)
    if err != nil {
        fmt.Println(err)
    }
    return &post, nil

}

func (s *BlogGRPCService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
    id := generateUniqueID()
    title := req.Title
    content := req.Content

    insert, err := s.DB.Query("INSERT INTO blog (id, title, content) VALUES (?, ?, ?)", id, title, content)
    if err != nil {
        fmt.Println(err)
    }

    defer insert.Close()

    return &pb.Post{
        Id: id,
        Title:   title,
        Content: content,
    }, nil

}

func (s *BlogGRPCService) DeletePost(ctx context.Context, req *pb.GetPostRequest) (*pb.Empty, error) {
    id := req.Id

    delete, err := s.DB.Query("DELETE FROM blog WHERE id = ?", id);
    if err != nil {
        fmt.Println(err)
    }

    defer delete.Close()

    return &pb.Empty{}, nil
}

func (s *BlogGRPCService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
    id := req.Id
    title := req.Title
    content := req.Content

    update, err := s.DB.Query("UPDATE blog SET title = ?, content = ? WHERE id = ?", title, content, id)
    if err != nil {
        fmt.Println(err)
    }

    defer update.Close()

    return &pb.Post{
        Id: id,
        Title:   title,
        Content: content,
    }, nil
}

func generateUniqueID() string {
	return uuid.New().String()
}
