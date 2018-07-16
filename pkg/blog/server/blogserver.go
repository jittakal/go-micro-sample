package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jittakal/go-micro-sample/pkg/blog/db"
	pb "github.com/jittakal/go-micro-sample/pkg/blog/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type (
	userServer    struct{}
	articleServer struct{}
)

func (s *userServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := db.User{
		Name:  in.Name,
		Email: in.Email,
	}

	id, err := user.Create()
	return &pb.CreateUserResponse{Id: id}, err
}

func (s *articleServer) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	fmt.Println("Title of the article : ", in.Title)
	fmt.Println("Content of the article : ", in.Content)
	return &pb.CreateArticleResponse{Id: "1234"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// Register user server
	pb.RegisterUserServer(s, &userServer{})
	// Register article server
	pb.RegisterArticleServer(s, &articleServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
