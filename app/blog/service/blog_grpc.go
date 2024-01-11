// service/blog_grpc.go
package service

import (
	"context"
	"fmt"
	"week1/app/blog/storage"
	"week1/app/blog/pb"
	"github.com/google/uuid"
	"errors"
)

type BlogGRPCService struct{
	pb.UnimplementedBlogServiceServer
}

func NewBlogGRPCService() *BlogGRPCService {
	return &BlogGRPCService{}
}

func (s *BlogGRPCService) GetPosts(ctx context.Context, req *pb.Empty) (*pb.PostList, error) {
    posts := storage.Posts
    protoPosts := make([]*pb.Post, len(posts))

    for i, p := range posts {
        protoPosts[i] = &pb.Post{
            Id:      p.ID,
            Title:   p.Title,
            Content: p.Content,
        }
    }

    return &pb.PostList{Posts: protoPosts}, nil
}

func (s *BlogGRPCService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
    id := req.Id
    fmt.Println("GetPostByID", id)
    for _, p := range storage.Posts {
        if p.ID == id {
            return &pb.Post{
                Id:      p.ID,
                Title:   p.Title,
                Content: p.Content,
            }, nil
        }
    }
    return nil, errors.New("post not found")
}

func (s *BlogGRPCService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
    newPost := storage.Post{
        ID:      generateUniqueID(),
        Title:   req.Title,
        Content: req.Content,
    }
    storage.Posts = append(storage.Posts, newPost)

    // Return the protobuf equivalent of the new post
    return &pb.Post{
        Id:      newPost.ID,
        Title:   newPost.Title,
        Content: newPost.Content,
    }, nil
}

func (s *BlogGRPCService) DeletePost(ctx context.Context, req *pb.GetPostRequest) (*pb.Empty, error) {
	id := req.Id

	for i, p := range storage.Posts {
		if p.ID == id {
			storage.Posts = append(storage.Posts[:i], storage.Posts[i+1:]...)
			break
		}
	}

	return &pb.Empty{}, nil
}

func (s *BlogGRPCService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	id := req.Id
	fmt.Println("UpdatePostByID", req)
	fmt.Println("id", id)

	for i, p := range storage.Posts {
		if p.ID == id {
			storage.Posts[i].Title = req.Title
			storage.Posts[i].Content = req.Content
			break
		}
	}

	return &pb.Post{
		Id:      id,
		Title:   req.Title,
		Content: req.Content,
	}, nil
}

func generateUniqueID() string {
	return uuid.New().String()
}
