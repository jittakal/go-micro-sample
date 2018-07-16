package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jittakal/go-micro-sample/pkg/blog/service"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "Arnav Takalkar",
		Email: "arnav.takalkar@gmail.com",
	})
	if err != nil {
		log.Fatalf("could not reach to user service server: %v", err)
	}
	log.Printf("New User Id : %s", r.Id)

	ac := pb.NewArticleClient(conn)
	ar, err := ac.CreateArticle(ctx, &pb.CreateArticleRequest{
		Title:   "My First Blog",
		Content: "Hello World!",
	})
	log.Printf("New Article Id : %s", ar.Id)

}
