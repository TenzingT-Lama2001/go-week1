// service/blog_grpc.go
package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "week1/app/gen"

	"github.com/google/uuid"
)

type BlogGRPCService struct {
	pb.UnimplementedBlogServiceServer
	DB *sql.DB
}

func NewBlogGRPCService(db *sql.DB) *BlogGRPCService {
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
		Id:      id,
		Title:   title,
		Content: content,
	}, nil

}

func (s *BlogGRPCService) DeletePost(ctx context.Context, req *pb.GetPostRequest) (*pb.Empty, error) {
	id := req.Id

	delete, err := s.DB.Query("DELETE FROM blog WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
	}

	defer delete.Close()

	return &pb.Empty{}, nil
}

func (s *BlogGRPCService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	id := req.Id
	var title string
	var content string
	query := "UPDATE blog SET "
	params := []interface{}{}

	if req.Title != nil {
		title = *req.Title
		query += "title = ?, "
		params = append(params, title)
	}

	if req.Content != nil {
		content = *req.Content
		query += "content = ? "
		params = append(params, content)

	}

	if req.Title == nil && req.Content == nil {
		return nil, errors.New("title or Content must be provided")
	}

	query += " WHERE id = ?"
	params = append(params, id)

	fmt.Println("query", query)
	fmt.Println("params", params)
	update, err := s.DB.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
	}

	affectedRows, err := update.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to check if post exists")
	}

	if affectedRows == 0 {
		return nil, errors.New("no rows affected. Either post is already updated or post does not exist")
	}

	return &pb.Post{
		Id:      id,
		Title:   title,
		Content: content,
	}, nil
}

func (s *BlogGRPCService) PostExists(ctx context.Context, req *pb.PostExistsRequest) (*pb.PostExistsResponse, error) {
	title := req.Title
	fmt.Println("title", title)
	blog := s.DB.QueryRow("SELECT * FROM blog WHERE title = ?", title)

	//if blog exists return true
	var post pb.Post
	err := blog.Scan(&post.Id, &post.Title, &post.Content)
	if err != nil {
		fmt.Println(err)
		return &pb.PostExistsResponse{Exists: false}, nil
	}
	return &pb.PostExistsResponse{Exists: true}, nil
}

func generateUniqueID() string {
	return uuid.New().String()
}
